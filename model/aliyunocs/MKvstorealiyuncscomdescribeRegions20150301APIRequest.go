package aliyunocs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// MKvstorealiyuncscomdescribeRegions20150301APIRequest 查看Region列表 API请求
// m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01
//
// 查看Region列表
type MKvstorealiyuncscomdescribeRegions20150301APIRequest struct {
	model.Params
}

// NewMKvstorealiyuncscomdescribeRegions20150301Request 初始化MKvstorealiyuncscomdescribeRegions20150301APIRequest对象
func NewMKvstorealiyuncscomdescribeRegions20150301Request() *MKvstorealiyuncscomdescribeRegions20150301APIRequest {
	return &MKvstorealiyuncscomdescribeRegions20150301APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r MKvstorealiyuncscomdescribeRegions20150301APIRequest) GetApiMethodName() string {
	return "m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r MKvstorealiyuncscomdescribeRegions20150301APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r MKvstorealiyuncscomdescribeRegions20150301APIRequest) GetRawParams() model.Params {
	return r.Params
}
