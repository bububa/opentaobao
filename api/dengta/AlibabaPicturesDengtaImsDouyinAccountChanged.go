package dengta

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/dengta"
)

/* 
接收发生变化的抖音帐号 
alibaba.pictures.dengta.ims.douyin.account.changed

接收发生变化的抖音帐号
*/
func AlibabaPicturesDengtaImsDouyinAccountChanged(clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaImsDouyinAccountChangedRequest, session string) (*dengta.AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse, error) {
    var resp dengta.AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}