package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gsp图片上传 APIRequest
alibaba.gsp.supply.image.upload

上传图片至目标海外平台的素材空间
*/
type AlibabaGspSupplyImageUploadRequest struct {
    model.Params

    // 图片名称
    fileName   string 

    // 图片文件流，像素宽度不小于500，不大于2000，像素长度不小于500，不大于2000
    fileContent   []*model.File 

}

func NewAlibabaGspSupplyImageUploadRequest() *AlibabaGspSupplyImageUploadRequest{
    return &AlibabaGspSupplyImageUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaGspSupplyImageUploadRequest) GetApiMethodName() string {
    return "alibaba.gsp.supply.image.upload"
}

func (r AlibabaGspSupplyImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaGspSupplyImageUploadRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

func (r AlibabaGspSupplyImageUploadRequest) GetFileName() string {
    return r.fileName
}

func (r *AlibabaGspSupplyImageUploadRequest) SetFileContent(fileContent []*model.File) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

func (r AlibabaGspSupplyImageUploadRequest) GetFileContent() []*model.File {
    return r.fileContent
}

