package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcarcarefreedetailgetAPIResponse 查询业务单信息 API返回值
// tmall.car.carefree.detail.get
//
// 查询业务单信息
type TmallcarcarefreedetailgetAPIResponse struct {
	model.CommonResponse
	TmallcarcarefreedetailgetAPIResponseModel
}

// TmallcarcarefreedetailgetAPIResponseModel is 查询业务单信息 成功返回结果
type TmallcarcarefreedetailgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_carefree_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误编号
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 数据对象
	Data *CarefreeDetailInfoDto `json:"data,omitempty" xml:"data,omitempty"`
}
