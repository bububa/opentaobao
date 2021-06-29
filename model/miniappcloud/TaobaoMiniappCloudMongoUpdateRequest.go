package miniappcloud

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新MongoDB中的数据 API请求
taobao.miniapp.cloud.mongo.update

更新MongoDB中的数据
*/
type TaobaoMiniappCloudMongoUpdateRequest struct {
    model.Params
    // MongoDB表名
    collection   string
    // 更新条件
    filter   string
    // 待写入的数据
    record   string
    // 要操作的环境，默认是测试环境
    env   string
}

// 初始化TaobaoMiniappCloudMongoUpdateRequest对象
func NewTaobaoMiniappCloudMongoUpdateRequest() *TaobaoMiniappCloudMongoUpdateRequest{
    return &TaobaoMiniappCloudMongoUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudMongoUpdateRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.mongo.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudMongoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Collection Setter
// MongoDB表名
func (r *TaobaoMiniappCloudMongoUpdateRequest) SetCollection(collection string) error {
    r.collection = collection
    r.Set("collection", collection)
    return nil
}

// Collection Getter
func (r TaobaoMiniappCloudMongoUpdateRequest) GetCollection() string {
    return r.collection
}
// Filter Setter
// 更新条件
func (r *TaobaoMiniappCloudMongoUpdateRequest) SetFilter(filter string) error {
    r.filter = filter
    r.Set("filter", filter)
    return nil
}

// Filter Getter
func (r TaobaoMiniappCloudMongoUpdateRequest) GetFilter() string {
    return r.filter
}
// Record Setter
// 待写入的数据
func (r *TaobaoMiniappCloudMongoUpdateRequest) SetRecord(record string) error {
    r.record = record
    r.Set("record", record)
    return nil
}

// Record Getter
func (r TaobaoMiniappCloudMongoUpdateRequest) GetRecord() string {
    return r.record
}
// Env Setter
// 要操作的环境，默认是测试环境
func (r *TaobaoMiniappCloudMongoUpdateRequest) SetEnv(env string) error {
    r.env = env
    r.Set("env", env)
    return nil
}

// Env Getter
func (r TaobaoMiniappCloudMongoUpdateRequest) GetEnv() string {
    return r.env
}
