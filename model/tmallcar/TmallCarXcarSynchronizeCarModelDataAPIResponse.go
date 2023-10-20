package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarXcarSynchronizeCarModelDataAPIResponse 爱车车型数据同步 API返回值
// tmall.car.xcar.synchronize.car.model.data
//
// 爱车汽车车型数据同步到天猫
type TmallCarXcarSynchronizeCarModelDataAPIResponse struct {
	model.CommonResponse
	TmallCarXcarSynchronizeCarModelDataAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarXcarSynchronizeCarModelDataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarXcarSynchronizeCarModelDataAPIResponseModel).Reset()
}

// TmallCarXcarSynchronizeCarModelDataAPIResponseModel is 爱车车型数据同步 成功返回结果
type TmallCarXcarSynchronizeCarModelDataAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_xcar_synchronize_car_model_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象描述
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarXcarSynchronizeCarModelDataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarXcarSynchronizeCarModelDataAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarXcarSynchronizeCarModelDataAPIResponse)
	},
}

// GetTmallCarXcarSynchronizeCarModelDataAPIResponse 从 sync.Pool 获取 TmallCarXcarSynchronizeCarModelDataAPIResponse
func GetTmallCarXcarSynchronizeCarModelDataAPIResponse() *TmallCarXcarSynchronizeCarModelDataAPIResponse {
	return poolTmallCarXcarSynchronizeCarModelDataAPIResponse.Get().(*TmallCarXcarSynchronizeCarModelDataAPIResponse)
}

// ReleaseTmallCarXcarSynchronizeCarModelDataAPIResponse 将 TmallCarXcarSynchronizeCarModelDataAPIResponse 保存到 sync.Pool
func ReleaseTmallCarXcarSynchronizeCarModelDataAPIResponse(v *TmallCarXcarSynchronizeCarModelDataAPIResponse) {
	v.Reset()
	poolTmallCarXcarSynchronizeCarModelDataAPIResponse.Put(v)
}
