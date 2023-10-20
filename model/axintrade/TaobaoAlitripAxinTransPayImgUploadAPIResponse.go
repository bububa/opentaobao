package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintranspayimguploadAPIResponse 上传图片到支付宝图片空间接口 API返回值
// taobao.alitrip.axin.trans.pay.img.upload
//
// 阿信供销平台-上传图片到支付宝图片空间接口
type TaobaoalitripaxintranspayimguploadAPIResponse struct {
	model.CommonResponse
	TaobaoalitripaxintranspayimguploadAPIResponseModel
}

// TaobaoalitripaxintranspayimguploadAPIResponseModel is 上传图片到支付宝图片空间接口 成功返回结果
type TaobaoalitripaxintranspayimguploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_pay_img_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoalitripaxintranspayimguploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
