package nrpos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoscommdyofflinegetfileurlAPIResponse 去前置机pos商品离线文件下载地址查询接口 API返回值
// alibaba.mos.commdy.offline.getfileurl
//
// 去前置机-pos查询离线文件下载地址接口
type AlibabamoscommdyofflinegetfileurlAPIResponse struct {
	model.CommonResponse
	AlibabamoscommdyofflinegetfileurlAPIResponseModel
}

// AlibabamoscommdyofflinegetfileurlAPIResponseModel is 去前置机pos商品离线文件下载地址查询接口 成功返回结果
type AlibabamoscommdyofflinegetfileurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_commdy_offline_getfileurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabamoscommdyofflinegetfileurlResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
