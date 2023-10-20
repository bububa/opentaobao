package aliyunocs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// MKvstoreAliyuncsComDescribeRegions20150301APIRequest 查看Region列表 API请求
// m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01
//
// 查看Region列表
type MKvstoreAliyuncsComDescribeRegions20150301APIRequest struct {
	model.Params
}

// NewMKvstoreAliyuncsComDescribeRegions20150301Request 初始化MKvstoreAliyuncsComDescribeRegions20150301APIRequest对象
func NewMKvstoreAliyuncsComDescribeRegions20150301Request() *MKvstoreAliyuncsComDescribeRegions20150301APIRequest {
	return &MKvstoreAliyuncsComDescribeRegions20150301APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r MKvstoreAliyuncsComDescribeRegions20150301APIRequest) GetApiMethodName() string {
	return "m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r MKvstoreAliyuncsComDescribeRegions20150301APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r MKvstoreAliyuncsComDescribeRegions20150301APIRequest) GetRawParams() model.Params {
	return r.Params
}
