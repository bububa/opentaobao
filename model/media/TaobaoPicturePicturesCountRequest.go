package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片总数查询 APIRequest
taobao.picture.pictures.count

图片总数查询
*/
type TaobaoPicturePicturesCountRequest struct {
    model.Params

    // 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
    startDate   string 

    // 图片分类
    pictureCategoryId   int64 

    // 文件名
    title   string 

    // 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
    endDate   string 

    // 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
    startModifiedDate   string 

    // 是否删除,undeleted代表没有删除,deleted表示删除
    deleted   string 

    // 图片ID
    pictureId   int64 

    // 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
    clientType   string 

}

func NewTaobaoPicturePicturesCountRequest() *TaobaoPicturePicturesCountRequest{
    return &TaobaoPicturePicturesCountRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPicturePicturesCountRequest) GetApiMethodName() string {
    return "taobao.picture.pictures.count"
}

func (r TaobaoPicturePicturesCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPicturePicturesCountRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoPicturePicturesCountRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoPicturePicturesCountRequest) SetPictureCategoryId(pictureCategoryId int64) error {
    r.pictureCategoryId = pictureCategoryId
    r.Set("picture_category_id", pictureCategoryId)
    return nil
}

func (r TaobaoPicturePicturesCountRequest) GetPictureCategoryId() int64 {
    return r.pictureCategoryId
}

func (r *TaobaoPicturePicturesCountRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoPicturePicturesCountRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoPicturePicturesCountRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoPicturePicturesCountRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoPicturePicturesCountRequest) SetStartModifiedDate(startModifiedDate string) error {
    r.startModifiedDate = startModifiedDate
    r.Set("start_modified_date", startModifiedDate)
    return nil
}

func (r TaobaoPicturePicturesCountRequest) GetStartModifiedDate() string {
    return r.startModifiedDate
}

func (r *TaobaoPicturePicturesCountRequest) SetDeleted(deleted string) error {
    r.deleted = deleted
    r.Set("deleted", deleted)
    return nil
}

func (r TaobaoPicturePicturesCountRequest) GetDeleted() string {
    return r.deleted
}

func (r *TaobaoPicturePicturesCountRequest) SetPictureId(pictureId int64) error {
    r.pictureId = pictureId
    r.Set("picture_id", pictureId)
    return nil
}

func (r TaobaoPicturePicturesCountRequest) GetPictureId() int64 {
    return r.pictureId
}

func (r *TaobaoPicturePicturesCountRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r TaobaoPicturePicturesCountRequest) GetClientType() string {
    return r.clientType
}

