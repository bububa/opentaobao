package aliyun

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// OdpsAliyuncsComDescribeOdpsService20141215APIRequest 查询ODPS服务 API请求
// odps.aliyuncs.com.DescribeOdpsService.2014-12-15
//
// 查询ODPS服务
type OdpsAliyuncsComDescribeOdpsService20141215APIRequest struct {
	model.Params
}

// NewOdpsAliyuncsComDescribeOdpsService20141215Request 初始化OdpsAliyuncsComDescribeOdpsService20141215APIRequest对象
func NewOdpsAliyuncsComDescribeOdpsService20141215Request() *OdpsAliyuncsComDescribeOdpsService20141215APIRequest {
	return &OdpsAliyuncsComDescribeOdpsService20141215APIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *OdpsAliyuncsComDescribeOdpsService20141215APIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r OdpsAliyuncsComDescribeOdpsService20141215APIRequest) GetApiMethodName() string {
	return "odps.aliyuncs.com.DescribeOdpsService.2014-12-15"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r OdpsAliyuncsComDescribeOdpsService20141215APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r OdpsAliyuncsComDescribeOdpsService20141215APIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolOdpsAliyuncsComDescribeOdpsService20141215APIRequest = sync.Pool{
	New: func() any {
		return NewOdpsAliyuncsComDescribeOdpsService20141215Request()
	},
}

// GetOdpsAliyuncsComDescribeOdpsService20141215Request 从 sync.Pool 获取 OdpsAliyuncsComDescribeOdpsService20141215APIRequest
func GetOdpsAliyuncsComDescribeOdpsService20141215APIRequest() *OdpsAliyuncsComDescribeOdpsService20141215APIRequest {
	return poolOdpsAliyuncsComDescribeOdpsService20141215APIRequest.Get().(*OdpsAliyuncsComDescribeOdpsService20141215APIRequest)
}

// ReleaseOdpsAliyuncsComDescribeOdpsService20141215APIRequest 将 OdpsAliyuncsComDescribeOdpsService20141215APIRequest 放入 sync.Pool
func ReleaseOdpsAliyuncsComDescribeOdpsService20141215APIRequest(v *OdpsAliyuncsComDescribeOdpsService20141215APIRequest) {
	v.Reset()
	poolOdpsAliyuncsComDescribeOdpsService20141215APIRequest.Put(v)
}
