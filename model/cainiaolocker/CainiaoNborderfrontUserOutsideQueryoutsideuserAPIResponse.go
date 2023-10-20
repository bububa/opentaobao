package cainiaolocker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse 查询外部小件员休息 API返回值
// cainiao.nborderfront.user.outside.queryoutsideuser
//
// 采用SPI方式查询外部公司的小件员信息
type CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse struct {
	model.CommonResponse
	CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponseModel).Reset()
}

// CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponseModel is 查询外部小件员休息 成功返回结果
type CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_nborderfront_user_outside_queryoutsideuser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回错误编码
	RespCode string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// userInfo
	UserInfo *CainiaoNborderfrontUserOutsideQueryoutsideuserStruct `json:"user_info,omitempty" xml:"user_info,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespCode = ""
	m.UserInfo = nil
}

var poolCainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse)
	},
}

// GetCainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse 从 sync.Pool 获取 CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse
func GetCainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse() *CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse {
	return poolCainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse.Get().(*CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse)
}

// ReleaseCainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse 将 CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse 保存到 sync.Pool
func ReleaseCainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse(v *CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse) {
	v.Reset()
	poolCainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse.Put(v)
}
