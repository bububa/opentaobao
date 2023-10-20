package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcitycoordinatesbatchuploadAPIResponse 上传信息计算飞猪国际城市 API返回值
// taobao.xhotel.city.coordinates.batch.upload
//
// 给供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程。目前只适用于国际城市，国内+港澳台暂不支持。
// 非实时计算接口，每次批量上传不少于1条的数据，后端离线计算，请于30分钟后再下载结果。
type TaobaoxhotelcitycoordinatesbatchuploadAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelcitycoordinatesbatchuploadAPIResponseModel
}

// TaobaoxhotelcitycoordinatesbatchuploadAPIResponseModel is 上传信息计算飞猪国际城市 成功返回结果
type TaobaoxhotelcitycoordinatesbatchuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_city_coordinates_batch_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传成功后的批次号
	BatchId int64 `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
}
