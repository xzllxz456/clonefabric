/*
 * SPDX-License-Identifier: Apache-2.0
 */

package org.example;

import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;
import com.owlike.genson.Genson;

@DataType()
public class Java {

    private final static Genson genson = new Genson();

    @Property()
    private String value;

    public Java(){
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }

    public String toJSONString() {
        return genson.serialize(this).toString();
    }

    public static Java fromJSONString(String json) {
        Java asset = genson.deserialize(json, Java.class);
        return asset;
    }
}
