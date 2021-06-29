package scs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里妈妈智能创意平台在线抠图 APIRequest
alibaba.scs.image.matte

该API对外输出一个在线抠图(Deep Image Matting)接口，合作方可以通过该接口利用深度学习抠图算法从图片中抠出目标对象(比如商品或者人物轮廓)
*/
type AlibabaScsImageMatteRequest struct {
    model.Params

    // 资源位ID，接入前由智能创意平台分配
    pid   string 

    // 服务名称，可选值: scs
    name   string 

    // 场景名称，可选值: image_cutout
    scenes   string 

    // 抠图上下文信息，json字符串格式，json中matting_type字段可选值: external_matting，url: 需要抠图的目标图片url
    obExt   string 

    // 32位uuid
    sessionid   string 

    // 当前秒级时间戳
    ts   string 

}

func NewAlibabaScsImageMatteRequest() *AlibabaScsImageMatteRequest{
    return &AlibabaScsImageMatteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScsImageMatteRequest) GetApiMethodName() string {
    return "alibaba.scs.image.matte"
}

func (r AlibabaScsImageMatteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScsImageMatteRequest) SetPid(pid string) error {
    r.pid = pid
    r.Set("pid", pid)
    return nil
}

func (r AlibabaScsImageMatteRequest) GetPid() string {
    return r.pid
}

func (r *AlibabaScsImageMatteRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaScsImageMatteRequest) GetName() string {
    return r.name
}

func (r *AlibabaScsImageMatteRequest) SetScenes(scenes string) error {
    r.scenes = scenes
    r.Set("scenes", scenes)
    return nil
}

func (r AlibabaScsImageMatteRequest) GetScenes() string {
    return r.scenes
}

func (r *AlibabaScsImageMatteRequest) SetObExt(obExt string) error {
    r.obExt = obExt
    r.Set("ob_ext", obExt)
    return nil
}

func (r AlibabaScsImageMatteRequest) GetObExt() string {
    return r.obExt
}

func (r *AlibabaScsImageMatteRequest) SetSessionid(sessionid string) error {
    r.sessionid = sessionid
    r.Set("sessionid", sessionid)
    return nil
}

func (r AlibabaScsImageMatteRequest) GetSessionid() string {
    return r.sessionid
}

func (r *AlibabaScsImageMatteRequest) SetTs(ts string) error {
    r.ts = ts
    r.Set("ts", ts)
    return nil
}

func (r AlibabaScsImageMatteRequest) GetTs() string {
    return r.ts
}

