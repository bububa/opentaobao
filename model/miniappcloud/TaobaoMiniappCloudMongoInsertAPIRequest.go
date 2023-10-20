package miniappcloud

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudMongoInsertAPIRequest MongoDB插入单条数据 API请求
// taobao.miniapp.cloud.mongo.insert
//
// 向商家应用云中插入一条记录，用于外部数据同步到应用中
type TaobaoMiniappCloudMongoInsertAPIRequest struct {
	model.Params
	// 待插入的数据，JSON格式
	_record string
	// MongoDB表名
	_collection string
	// 要操作的环境，默认是测试环境
	_env string
}

// NewTaobaoMiniappCloudMongoInsertRequest 初始化TaobaoMiniappCloudMongoInsertAPIRequest对象
func NewTaobaoMiniappCloudMongoInsertRequest() *TaobaoMiniappCloudMongoInsertAPIRequest {
	return &TaobaoMiniappCloudMongoInsertAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappCloudMongoInsertAPIRequest) Reset() {
	r._record = ""
	r._collection = ""
	r._env = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudMongoInsertAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.mongo.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudMongoInsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappCloudMongoInsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecord is Record Setter
// 待插入的数据，JSON格式
func (r *TaobaoMiniappCloudMongoInsertAPIRequest) SetRecord(_record string) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// GetRecord Record Getter
func (r TaobaoMiniappCloudMongoInsertAPIRequest) GetRecord() string {
	return r._record
}

// SetCollection is Collection Setter
// MongoDB表名
func (r *TaobaoMiniappCloudMongoInsertAPIRequest) SetCollection(_collection string) error {
	r._collection = _collection
	r.Set("collection", _collection)
	return nil
}

// GetCollection Collection Getter
func (r TaobaoMiniappCloudMongoInsertAPIRequest) GetCollection() string {
	return r._collection
}

// SetEnv is Env Setter
// 要操作的环境，默认是测试环境
func (r *TaobaoMiniappCloudMongoInsertAPIRequest) SetEnv(_env string) error {
	r._env = _env
	r.Set("env", _env)
	return nil
}

// GetEnv Env Getter
func (r TaobaoMiniappCloudMongoInsertAPIRequest) GetEnv() string {
	return r._env
}

var poolTaobaoMiniappCloudMongoInsertAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappCloudMongoInsertRequest()
	},
}

// GetTaobaoMiniappCloudMongoInsertRequest 从 sync.Pool 获取 TaobaoMiniappCloudMongoInsertAPIRequest
func GetTaobaoMiniappCloudMongoInsertAPIRequest() *TaobaoMiniappCloudMongoInsertAPIRequest {
	return poolTaobaoMiniappCloudMongoInsertAPIRequest.Get().(*TaobaoMiniappCloudMongoInsertAPIRequest)
}

// ReleaseTaobaoMiniappCloudMongoInsertAPIRequest 将 TaobaoMiniappCloudMongoInsertAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappCloudMongoInsertAPIRequest(v *TaobaoMiniappCloudMongoInsertAPIRequest) {
	v.Reset()
	poolTaobaoMiniappCloudMongoInsertAPIRequest.Put(v)
}
