package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// OdpsaliyuncscomdescribeOdpsService20141215APIRequest 查询ODPS服务 API请求
// odps.aliyuncs.com.DescribeOdpsService.2014-12-15
//
// 查询ODPS服务
type OdpsaliyuncscomdescribeOdpsService20141215APIRequest struct {
	model.Params
}

// NewOdpsaliyuncscomdescribeOdpsService20141215Request 初始化OdpsaliyuncscomdescribeOdpsService20141215APIRequest对象
func NewOdpsaliyuncscomdescribeOdpsService20141215Request() *OdpsaliyuncscomdescribeOdpsService20141215APIRequest {
	return &OdpsaliyuncscomdescribeOdpsService20141215APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r OdpsaliyuncscomdescribeOdpsService20141215APIRequest) GetApiMethodName() string {
	return "odps.aliyuncs.com.DescribeOdpsService.2014-12-15"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r OdpsaliyuncscomdescribeOdpsService20141215APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r OdpsaliyuncscomdescribeOdpsService20141215APIRequest) GetRawParams() model.Params {
	return r.Params
}
