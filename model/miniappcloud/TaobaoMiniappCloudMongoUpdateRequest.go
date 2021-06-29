package miniappcloud

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新MongoDB中的数据 APIRequest
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

func NewTaobaoMiniappCloudMongoUpdateRequest() *TaobaoMiniappCloudMongoUpdateRequest{
    return &TaobaoMiniappCloudMongoUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappCloudMongoUpdateRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.mongo.update"
}

func (r TaobaoMiniappCloudMongoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappCloudMongoUpdateRequest) SetCollection(collection string) error {
    r.collection = collection
    r.Set("collection", collection)
    return nil
}

func (r TaobaoMiniappCloudMongoUpdateRequest) GetCollection() string {
    return r.collection
}

func (r *TaobaoMiniappCloudMongoUpdateRequest) SetFilter(filter string) error {
    r.filter = filter
    r.Set("filter", filter)
    return nil
}

func (r TaobaoMiniappCloudMongoUpdateRequest) GetFilter() string {
    return r.filter
}

func (r *TaobaoMiniappCloudMongoUpdateRequest) SetRecord(record string) error {
    r.record = record
    r.Set("record", record)
    return nil
}

func (r TaobaoMiniappCloudMongoUpdateRequest) GetRecord() string {
    return r.record
}

func (r *TaobaoMiniappCloudMongoUpdateRequest) SetEnv(env string) error {
    r.env = env
    r.Set("env", env)
    return nil
}

func (r TaobaoMiniappCloudMongoUpdateRequest) GetEnv() string {
    return r.env
}

