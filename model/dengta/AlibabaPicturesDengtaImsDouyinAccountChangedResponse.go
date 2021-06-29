package dengta

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
接收发生变化的抖音帐号 APIResponse
alibaba.pictures.dengta.ims.douyin.account.changed

接收发生变化的抖音帐号
*/
type AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse struct {
    model.CommonResponse
    AlibabaPicturesDengtaImsDouyinAccountChangedResponse
}

type AlibabaPicturesDengtaImsDouyinAccountChangedResponse struct {
    XMLName xml.Name `xml:"alibaba_pictures_dengta_ims_douyin_account_changed_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *GeneralResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
