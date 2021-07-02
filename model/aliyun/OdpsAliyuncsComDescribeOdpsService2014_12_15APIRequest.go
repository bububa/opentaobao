package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// OdpsAliyuncsComDescribeOdpsService2014_12_15APIRequest 查询ODPS服务 API请求
// odps.aliyuncs.com.DescribeOdpsService.2014-12-15
//
// 查询ODPS服务
type OdpsAliyuncsComDescribeOdpsService2014_12_15APIRequest struct {
	model.Params
}

// NewOdpsAliyuncsComDescribeOdpsService2014_12_15Request 初始化OdpsAliyuncsComDescribeOdpsService2014_12_15APIRequest对象
func NewOdpsAliyuncsComDescribeOdpsService2014_12_15Request() *OdpsAliyuncsComDescribeOdpsService2014_12_15APIRequest {
	return &OdpsAliyuncsComDescribeOdpsService2014_12_15APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r OdpsAliyuncsComDescribeOdpsService2014_12_15APIRequest) GetApiMethodName() string {
	return "odps.aliyuncs.com.DescribeOdpsService.2014-12-15"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r OdpsAliyuncsComDescribeOdpsService2014_12_15APIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
