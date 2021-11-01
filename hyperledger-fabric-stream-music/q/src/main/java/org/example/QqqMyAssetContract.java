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

@Contract(name = "QqqMyAssetContract",
    info = @Info(title = "QqqMyAsset contract",
                description = "My Smart Contract",
                version = "0.0.1",
                license =
                        @License(name = "Apache-2.0",
                                url = ""),
                                contact =  @Contact(email = "q@example.com",
                                                name = "q",
                                                url = "http://q.me")))
@Default
public class QqqMyAssetContract implements ContractInterface {
    public  QqqMyAssetContract() {

    }
    @Transaction()
    public boolean qqqMyAssetExists(Context ctx, String qqqMyAssetId) {
        byte[] buffer = ctx.getStub().getState(qqqMyAssetId);
        return (buffer != null && buffer.length > 0);
    }

    @Transaction()
    public void createQqqMyAsset(Context ctx, String qqqMyAssetId, String value) {
        boolean exists = qqqMyAssetExists(ctx,qqqMyAssetId);
        if (exists) {
            throw new RuntimeException("The asset "+qqqMyAssetId+" already exists");
        }
        QqqMyAsset asset = new QqqMyAsset();
        asset.setValue(value);
        ctx.getStub().putState(qqqMyAssetId, asset.toJSONString().getBytes(UTF_8));
    }

    @Transaction()
    public QqqMyAsset readQqqMyAsset(Context ctx, String qqqMyAssetId) {
        boolean exists = qqqMyAssetExists(ctx,qqqMyAssetId);
        if (!exists) {
            throw new RuntimeException("The asset "+qqqMyAssetId+" does not exist");
        }

        QqqMyAsset newAsset = QqqMyAsset.fromJSONString(new String(ctx.getStub().getState(qqqMyAssetId),UTF_8));
        return newAsset;
    }

    @Transaction()
    public void updateQqqMyAsset(Context ctx, String qqqMyAssetId, String newValue) {
        boolean exists = qqqMyAssetExists(ctx,qqqMyAssetId);
        if (!exists) {
            throw new RuntimeException("The asset "+qqqMyAssetId+" does not exist");
        }
        QqqMyAsset asset = new QqqMyAsset();
        asset.setValue(newValue);

        ctx.getStub().putState(qqqMyAssetId, asset.toJSONString().getBytes(UTF_8));
    }

    @Transaction()
    public void deleteQqqMyAsset(Context ctx, String qqqMyAssetId) {
        boolean exists = qqqMyAssetExists(ctx,qqqMyAssetId);
        if (!exists) {
            throw new RuntimeException("The asset "+qqqMyAssetId+" does not exist");
        }
        ctx.getStub().delState(qqqMyAssetId);
    }

}
