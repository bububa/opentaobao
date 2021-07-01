package alihealthlab

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthlab"
)

/* 
检验检测业务，isv项目门店关系同步 
alibaba.alihealth.lab.item.store.relation.sync

阿里健康检验检测业务，isv检验检测项目门店关系同步到健康，支持检验检测项目门店关系的增加和删除
*/
func AlibabaAlihealthLabItemStoreRelationSync(clt *core.SDKClient, req *alihealthlab.AlibabaAlihealthLabItemStoreRelationSyncAPIRequest, session string) (*alihealthlab.AlibabaAlihealthLabItemStoreRelationSyncAPIResponse, error) {
    var resp alihealthlab.AlibabaAlihealthLabItemStoreRelationSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
