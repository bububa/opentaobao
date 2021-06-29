package miniappcloud

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
MongoDB插入单条数据 API请求
taobao.miniapp.cloud.mongo.insert

向商家应用云中插入一条记录，用于外部数据同步到应用中
*/
type TaobaoMiniappCloudMongoInsertRequest struct {
    model.Params
    // 待插入的数据，JSON格式
    _record   string
    // MongoDB表名
    _collection   string
    // 要操作的环境，默认是测试环境
    _env   string
}

// 初始化TaobaoMiniappCloudMongoInsertRequest对象
func NewTaobaoMiniappCloudMongoInsertRequest() *TaobaoMiniappCloudMongoInsertRequest{
    return &TaobaoMiniappCloudMongoInsertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudMongoInsertRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.mongo.insert"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudMongoInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Record Setter
// 待插入的数据，JSON格式
func (r *TaobaoMiniappCloudMongoInsertRequest) SetRecord(_record string) error {
    r._record = _record
    r.Set("record", _record)
    return nil
}

// Record Getter
func (r TaobaoMiniappCloudMongoInsertRequest) GetRecord() string {
    return r._record
}
// Collection Setter
// MongoDB表名
func (r *TaobaoMiniappCloudMongoInsertRequest) SetCollection(_collection string) error {
    r._collection = _collection
    r.Set("collection", _collection)
    return nil
}

// Collection Getter
func (r TaobaoMiniappCloudMongoInsertRequest) GetCollection() string {
    return r._collection
}
// Env Setter
// 要操作的环境，默认是测试环境
func (r *TaobaoMiniappCloudMongoInsertRequest) SetEnv(_env string) error {
    r._env = _env
    r.Set("env", _env)
    return nil
}

// Env Getter
func (r TaobaoMiniappCloudMongoInsertRequest) GetEnv() string {
    return r._env
}
