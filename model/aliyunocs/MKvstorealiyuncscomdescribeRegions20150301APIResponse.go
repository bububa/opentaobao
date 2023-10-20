package aliyunocs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// MKvstoreAliyuncsComDescribeRegions20150301APIResponse 查看Region列表 API返回值
// m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01
//
// 查看Region列表
type MKvstoreAliyuncsComDescribeRegions20150301APIResponse struct {
	model.CommonResponse
	MKvstoreAliyuncsComDescribeRegions20150301APIResponseModel
}

// Reset 清空结构体
func (m *MKvstoreAliyuncsComDescribeRegions20150301APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.MKvstoreAliyuncsComDescribeRegions20150301APIResponseModel).Reset()
}

// MKvstoreAliyuncsComDescribeRegions20150301APIResponseModel is 查看Region列表 成功返回结果
type MKvstoreAliyuncsComDescribeRegions20150301APIResponseModel struct {
	XMLName xml.Name `xml:"m-kvstore_aliyuncs_com_DescribeRegions_2015-03-01_response"`
	// RegionIds 为一个 List,里面每个元素由三部分组成【RegionId,LocalName,ZoneIds】，其中ZoneIds是多个值组成的String，多个值间用英文半角“,”分隔。
	RegionIds []AliyunOcsRegion `json:"RegionIds,omitempty" xml:"RegionIds>aliyun_ocs_region,omitempty"`
	// 请求的唯一ID
	RequestId string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

// Reset 清空结构体
func (m *MKvstoreAliyuncsComDescribeRegions20150301APIResponseModel) Reset() {
	m.RequestId = ""
	m.RegionIds = m.RegionIds[:0]
	m.RequestId = ""
}

var poolMKvstoreAliyuncsComDescribeRegions20150301APIResponse = sync.Pool{
	New: func() any {
		return new(MKvstoreAliyuncsComDescribeRegions20150301APIResponse)
	},
}

// GetMKvstoreAliyuncsComDescribeRegions20150301APIResponse 从 sync.Pool 获取 MKvstoreAliyuncsComDescribeRegions20150301APIResponse
func GetMKvstoreAliyuncsComDescribeRegions20150301APIResponse() *MKvstoreAliyuncsComDescribeRegions20150301APIResponse {
	return poolMKvstoreAliyuncsComDescribeRegions20150301APIResponse.Get().(*MKvstoreAliyuncsComDescribeRegions20150301APIResponse)
}

// ReleaseMKvstoreAliyuncsComDescribeRegions20150301APIResponse 将 MKvstoreAliyuncsComDescribeRegions20150301APIResponse 保存到 sync.Pool
func ReleaseMKvstoreAliyuncsComDescribeRegions20150301APIResponse(v *MKvstoreAliyuncsComDescribeRegions20150301APIResponse) {
	v.Reset()
	poolMKvstoreAliyuncsComDescribeRegions20150301APIResponse.Put(v)
}
