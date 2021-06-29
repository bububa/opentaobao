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
type AlibabaAilabsAligenieSkillEntityImportRequest struct {
    model.Params
    // 技能Id
    skillId   int64
    // 是否增量更新，true为增量，否则上传数据中没有的实体值将被删除
    append   bool
    // 测试，目前没有此功能，此参数无效
    test   bool
    // 文件内容
    fileData   []*model.File
    // 要更新的实体名
    entityName   string
    // 固定值，ISV
    type   string
}

// 初始化AlibabaAilabsAligenieSkillEntityImportRequest对象
func NewAlibabaAilabsAligenieSkillEntityImportRequest() *AlibabaAilabsAligenieSkillEntityImportRequest{
    return &AlibabaAilabsAligenieSkillEntityImportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.skill.entity.import"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkillId Setter
// 技能Id
func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

// SkillId Getter
func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetSkillId() int64 {
    return r.skillId
}
// Append Setter
// 是否增量更新，true为增量，否则上传数据中没有的实体值将被删除
func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetAppend(append bool) error {
    r.append = append
    r.Set("append", append)
    return nil
}

// Append Getter
func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetAppend() bool {
    return r.append
}
// Test Setter
// 测试，目前没有此功能，此参数无效
func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetTest(test bool) error {
    r.test = test
    r.Set("test", test)
    return nil
}

// Test Getter
func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetTest() bool {
    return r.test
}
// FileData Setter
// 文件内容
func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetFileData(fileData []*model.File) error {
    r.fileData = fileData
    r.Set("file_data", fileData)
    return nil
}

// FileData Getter
func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetFileData() []*model.File {
    return r.fileData
}
// EntityName Setter
// 要更新的实体名
func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetEntityName(entityName string) error {
    r.entityName = entityName
    r.Set("entity_name", entityName)
    return nil
}

// EntityName Getter
func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetEntityName() string {
    return r.entityName
}
// Type Setter
// 固定值，ISV
func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetType() string {
    return r.type
}
