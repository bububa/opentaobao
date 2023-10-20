package miniappcloud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappcloudmongoinsertAPIRequest MongoDB插入单条数据 API请求
// taobao.miniapp.cloud.mongo.insert
//
// 向商家应用云中插入一条记录，用于外部数据同步到应用中
type TaobaominiappcloudmongoinsertAPIRequest struct {
	model.Params
	// 待插入的数据，JSON格式
	_record string
	// MongoDB表名
	_collection string
	// 要操作的环境，默认是测试环境
	_env string
}

// NewTaobaominiappcloudmongoinsertRequest 初始化TaobaominiappcloudmongoinsertAPIRequest对象
func NewTaobaominiappcloudmongoinsertRequest() *TaobaominiappcloudmongoinsertAPIRequest {
	return &TaobaominiappcloudmongoinsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappcloudmongoinsertAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.mongo.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappcloudmongoinsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappcloudmongoinsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecord is Record Setter
// 待插入的数据，JSON格式
func (r *TaobaominiappcloudmongoinsertAPIRequest) SetRecord(_record string) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// GetRecord Record Getter
func (r TaobaominiappcloudmongoinsertAPIRequest) GetRecord() string {
	return r._record
}

// SetCollection is Collection Setter
// MongoDB表名
func (r *TaobaominiappcloudmongoinsertAPIRequest) SetCollection(_collection string) error {
	r._collection = _collection
	r.Set("collection", _collection)
	return nil
}

// GetCollection Collection Getter
func (r TaobaominiappcloudmongoinsertAPIRequest) GetCollection() string {
	return r._collection
}

// SetEnv is Env Setter
// 要操作的环境，默认是测试环境
func (r *TaobaominiappcloudmongoinsertAPIRequest) SetEnv(_env string) error {
	r._env = _env
	r.Set("env", _env)
	return nil
}

// GetEnv Env Getter
func (r TaobaominiappcloudmongoinsertAPIRequest) GetEnv() string {
	return r._env
}
