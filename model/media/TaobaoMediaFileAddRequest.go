package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多媒体平台文件添加 APIRequest
taobao.media.file.add

用户通过top上传文件到多媒体平台
*/
type TaobaoMediaFileAddRequest struct {
    model.Params

    // 文件属于的那个目录的目录编号
    dirId   int64 

    // 上传文件的名称
    name   string 

    // 额外信息
    ext   int64 

    // 文件
    fileData   []byte 

    // 接入多媒体平台的业务code
每个应用有一个特有的业务code
    bizCode   string 

}

func NewTaobaoMediaFileAddRequest() *TaobaoMediaFileAddRequest{
    return &TaobaoMediaFileAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMediaFileAddRequest) GetApiMethodName() string {
    return "taobao.media.file.add"
}

func (r TaobaoMediaFileAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMediaFileAddRequest) SetDirId(dirId int64) error {
    r.dirId = dirId
    r.Set("dir_id", dirId)
    return nil
}

func (r TaobaoMediaFileAddRequest) GetDirId() int64 {
    return r.dirId
}

func (r *TaobaoMediaFileAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoMediaFileAddRequest) GetName() string {
    return r.name
}

func (r *TaobaoMediaFileAddRequest) SetExt(ext int64) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoMediaFileAddRequest) GetExt() int64 {
    return r.ext
}

func (r *TaobaoMediaFileAddRequest) SetFileData(fileData []byte) error {
    r.fileData = fileData
    r.Set("file_data", fileData)
    return nil
}

func (r TaobaoMediaFileAddRequest) GetFileData() []byte {
    return r.fileData
}

func (r *TaobaoMediaFileAddRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r TaobaoMediaFileAddRequest) GetBizCode() string {
    return r.bizCode
}

