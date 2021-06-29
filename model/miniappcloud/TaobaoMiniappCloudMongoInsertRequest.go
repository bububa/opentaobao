package miniappcloud

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
MongoDB插入单条数据 APIRequest
taobao.miniapp.cloud.mongo.insert

向商家应用云中插入一条记录，用于外部数据同步到应用中
*/
type TaobaoMiniappCloudMongoInsertRequest struct {
    model.Params

    // 待插入的数据，JSON格式
    record   string 

    // MongoDB表名
    collection   string 

    // 要操作的环境，默认是测试环境
    env   string 

}

func NewTaobaoMiniappCloudMongoInsertRequest() *TaobaoMiniappCloudMongoInsertRequest{
    return &TaobaoMiniappCloudMongoInsertRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappCloudMongoInsertRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.mongo.insert"
}

func (r TaobaoMiniappCloudMongoInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappCloudMongoInsertRequest) SetRecord(record string) error {
    r.record = record
    r.Set("record", record)
    return nil
}

func (r TaobaoMiniappCloudMongoInsertRequest) GetRecord() string {
    return r.record
}

func (r *TaobaoMiniappCloudMongoInsertRequest) SetCollection(collection string) error {
    r.collection = collection
    r.Set("collection", collection)
    return nil
}

func (r TaobaoMiniappCloudMongoInsertRequest) GetCollection() string {
    return r.collection
}

func (r *TaobaoMiniappCloudMongoInsertRequest) SetEnv(env string) error {
    r.env = env
    r.Set("env", env)
    return nil
}

func (r TaobaoMiniappCloudMongoInsertRequest) GetEnv() string {
    return r.env
}

