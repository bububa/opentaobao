package dengta

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabapicturesdengtaimsdouyinaccountchangedAPIResponse 接收发生变化的抖音帐号 API返回值
// alibaba.pictures.dengta.ims.douyin.account.changed
//
// 接收发生变化的抖音帐号
type AlibabapicturesdengtaimsdouyinaccountchangedAPIResponse struct {
	model.CommonResponse
	AlibabapicturesdengtaimsdouyinaccountchangedAPIResponseModel
}

// AlibabapicturesdengtaimsdouyinaccountchangedAPIResponseModel is 接收发生变化的抖音帐号 成功返回结果
type AlibabapicturesdengtaimsdouyinaccountchangedAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pictures_dengta_ims_douyin_account_changed_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *GeneralResult `json:"result,omitempty" xml:"result,omitempty"`
}
