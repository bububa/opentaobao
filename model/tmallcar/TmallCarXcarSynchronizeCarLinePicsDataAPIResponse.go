package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarXcarSynchronizeCarLinePicsDataAPIResponse 爱卡车系图片数据接入 API返回值
// tmall.car.xcar.synchronize.car.line.pics.data
//
// 爱卡车系图片数据同步天猫汽车
type TmallCarXcarSynchronizeCarLinePicsDataAPIResponse struct {
	model.CommonResponse
	TmallCarXcarSynchronizeCarLinePicsDataAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarXcarSynchronizeCarLinePicsDataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarXcarSynchronizeCarLinePicsDataAPIResponseModel).Reset()
}

// TmallCarXcarSynchronizeCarLinePicsDataAPIResponseModel is 爱卡车系图片数据接入 成功返回结果
type TmallCarXcarSynchronizeCarLinePicsDataAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_xcar_synchronize_car_line_pics_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarXcarSynchronizeCarLinePicsDataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarXcarSynchronizeCarLinePicsDataAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarXcarSynchronizeCarLinePicsDataAPIResponse)
	},
}

// GetTmallCarXcarSynchronizeCarLinePicsDataAPIResponse 从 sync.Pool 获取 TmallCarXcarSynchronizeCarLinePicsDataAPIResponse
func GetTmallCarXcarSynchronizeCarLinePicsDataAPIResponse() *TmallCarXcarSynchronizeCarLinePicsDataAPIResponse {
	return poolTmallCarXcarSynchronizeCarLinePicsDataAPIResponse.Get().(*TmallCarXcarSynchronizeCarLinePicsDataAPIResponse)
}

// ReleaseTmallCarXcarSynchronizeCarLinePicsDataAPIResponse 将 TmallCarXcarSynchronizeCarLinePicsDataAPIResponse 保存到 sync.Pool
func ReleaseTmallCarXcarSynchronizeCarLinePicsDataAPIResponse(v *TmallCarXcarSynchronizeCarLinePicsDataAPIResponse) {
	v.Reset()
	poolTmallCarXcarSynchronizeCarLinePicsDataAPIResponse.Put(v)
}
