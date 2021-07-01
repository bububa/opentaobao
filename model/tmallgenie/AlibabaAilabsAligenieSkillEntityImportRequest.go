package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实体动态更新 API请求
alibaba.ailabs.aligenie.skill.entity.import

根据用户上传的实体信息，进行制定技能实体的动态变更
*/
type AlibabaAilabsAligenieSkillEntityImportAPIRequest struct {
    model.Params
    // 技能Id
    _skillId   int64
    // 是否增量更新，true为增量，否则上传数据中没有的实体值将被删除
    _append   bool
    // 测试，目前没有此功能，此参数无效
    _test   bool
    // 文件内容
    _fileData   *model.File
    // 要更新的实体名
    _entityName   string
    // 固定值，ISV
    _type   string
}

// 初始化AlibabaAilabsAligenieSkillEntityImportAPIRequest对象
func NewAlibabaAilabsAligenieSkillEntityImportRequest() *AlibabaAilabsAligenieSkillEntityImportAPIRequest{
    return &AlibabaAilabsAligenieSkillEntityImportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieSkillEntityImportAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.skill.entity.import"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieSkillEntityImportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkillId Setter
// 技能Id
func (r *AlibabaAilabsAligenieSkillEntityImportAPIRequest) SetSkillId(_skillId int64) error {
    r._skillId = _skillId
    r.Set("skill_id", _skillId)
    return nil
}

// SkillId Getter
func (r AlibabaAilabsAligenieSkillEntityImportAPIRequest) GetSkillId() int64 {
    return r._skillId
}
// Append Setter
// 是否增量更新，true为增量，否则上传数据中没有的实体值将被删除
func (r *AlibabaAilabsAligenieSkillEntityImportAPIRequest) SetAppend(_append bool) error {
    r._append = _append
    r.Set("append", _append)
    return nil
}

// Append Getter
func (r AlibabaAilabsAligenieSkillEntityImportAPIRequest) GetAppend() bool {
    return r._append
}
// Test Setter
// 测试，目前没有此功能，此参数无效
func (r *AlibabaAilabsAligenieSkillEntityImportAPIRequest) SetTest(_test bool) error {
    r._test = _test
    r.Set("test", _test)
    return nil
}

// Test Getter
func (r AlibabaAilabsAligenieSkillEntityImportAPIRequest) GetTest() bool {
    return r._test
}
// FileData Setter
// 文件内容
func (r *AlibabaAilabsAligenieSkillEntityImportAPIRequest) SetFileData(_fileData *model.File) error {
    r._fileData = _fileData
    r.Set("file_data", _fileData)
    return nil
}

// FileData Getter
func (r AlibabaAilabsAligenieSkillEntityImportAPIRequest) GetFileData() *model.File {
    return r._fileData
}
// EntityName Setter
// 要更新的实体名
func (r *AlibabaAilabsAligenieSkillEntityImportAPIRequest) SetEntityName(_entityName string) error {
    r._entityName = _entityName
    r.Set("entity_name", _entityName)
    return nil
}

// EntityName Getter
func (r AlibabaAilabsAligenieSkillEntityImportAPIRequest) GetEntityName() string {
    return r._entityName
}
// Type Setter
// 固定值，ISV
func (r *AlibabaAilabsAligenieSkillEntityImportAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAilabsAligenieSkillEntityImportAPIRequest) GetType() string {
    return r._type
}
