package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多媒体平台文件添加 API请求
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
    fileData   []*model.File
    // 接入多媒体平台的业务code每个应用有一个特有的业务code
    bizCode   string
}

// 初始化TaobaoMediaFileAddRequest对象
func NewTaobaoMediaFileAddRequest() *TaobaoMediaFileAddRequest{
    return &TaobaoMediaFileAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMediaFileAddRequest) GetApiMethodName() string {
    return "taobao.media.file.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMediaFileAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DirId Setter
// 文件属于的那个目录的目录编号
func (r *TaobaoMediaFileAddRequest) SetDirId(dirId int64) error {
    r.dirId = dirId
    r.Set("dir_id", dirId)
    return nil
}

// DirId Getter
func (r TaobaoMediaFileAddRequest) GetDirId() int64 {
    return r.dirId
}
// Name Setter
// 上传文件的名称
func (r *TaobaoMediaFileAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoMediaFileAddRequest) GetName() string {
    return r.name
}
// Ext Setter
// 额外信息
func (r *TaobaoMediaFileAddRequest) SetExt(ext int64) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoMediaFileAddRequest) GetExt() int64 {
    return r.ext
}
// FileData Setter
// 文件
func (r *TaobaoMediaFileAddRequest) SetFileData(fileData []*model.File) error {
    r.fileData = fileData
    r.Set("file_data", fileData)
    return nil
}

// FileData Getter
func (r TaobaoMediaFileAddRequest) GetFileData() []*model.File {
    return r.fileData
}
// BizCode Setter
// 接入多媒体平台的业务code每个应用有一个特有的业务code
func (r *TaobaoMediaFileAddRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

// BizCode Getter
func (r TaobaoMediaFileAddRequest) GetBizCode() string {
    return r.bizCode
}
