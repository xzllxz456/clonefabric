/*
 * SPDX-License-Identifier: Apache-2.0
 */
package org.example;

import org.hyperledger.fabric.contract.Context;
import org.hyperledger.fabric.contract.ContractInterface;
import org.hyperledger.fabric.contract.annotation.Contract;
import org.hyperledger.fabric.contract.annotation.Default;
import org.hyperledger.fabric.contract.annotation.Transaction;
import org.hyperledger.fabric.contract.annotation.Contact;
import org.hyperledger.fabric.contract.annotation.Info;
import org.hyperledger.fabric.contract.annotation.License;
import static java.nio.charset.StandardCharsets.UTF_8;

@Contract(name = "JavaContract",
    info = @Info(title = "Java contract",
                description = "My Smart Contract",
                version = "0.0.1",
                license =
                        @License(name = "Apache-2.0",
                                url = ""),
                                contact =  @Contact(email = "aa@example.com",
                                                name = "aa",
                                                url = "http://aa.me")))
@Default
public class JavaContract implements ContractInterface {
    public  JavaContract() {

    }
    @Transaction()
    public boolean javaExists(Context ctx, String javaId) {
        byte[] buffer = ctx.getStub().getState(javaId);
        return (buffer != null && buffer.length > 0);
    }

    @Transaction()
    public void createJava(Context ctx, String javaId, String value) {
        boolean exists = javaExists(ctx,javaId);
        if (exists) {
            throw new RuntimeException("The asset "+javaId+" already exists");
        }
        Java asset = new Java();
        asset.setValue(value);
        ctx.getStub().putState(javaId, asset.toJSONString().getBytes(UTF_8));
    }

    @Transaction()
    public Java readJava(Context ctx, String javaId) {
        boolean exists = javaExists(ctx,javaId);
        if (!exists) {
            throw new RuntimeException("The asset "+javaId+" does not exist");
        }

        Java newAsset = Java.fromJSONString(new String(ctx.getStub().getState(javaId),UTF_8));
        return newAsset;
    }

    @Transaction()
    public void updateJava(Context ctx, String javaId, String newValue) {
        boolean exists = javaExists(ctx,javaId);
        if (!exists) {
            throw new RuntimeException("The asset "+javaId+" does not exist");
        }
        Java asset = new Java();
        asset.setValue(newValue);

        ctx.getStub().putState(javaId, asset.toJSONString().getBytes(UTF_8));
    }

    @Transaction()
    public void deleteJava(Context ctx, String javaId) {
        boolean exists = javaExists(ctx,javaId);
        if (!exists) {
            throw new RuntimeException("The asset "+javaId+" does not exist");
        }
        ctx.getStub().delState(javaId);
    }

}
