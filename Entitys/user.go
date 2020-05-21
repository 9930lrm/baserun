package Entitys

import (
	"database/sql"
	"time"
)

type User struct {
	AccountStatus     sql.NullString  `gorm:"column:account_status" json:"account_status"`
	ArrivalDate       time.Time       `gorm:"column:arrival_date" json:"arrival_date"`
	BankAccount       sql.NullString  `gorm:"column:bank_account" json:"bank_account"`
	BankName          sql.NullString  `gorm:"column:bank_name" json:"bank_name"`
	BirthPlace        sql.NullString  `gorm:"column:birth_place" json:"birth_place"`
	Birthday          time.Time       `gorm:"column:birthday" json:"birthday"`
	ContractPeriod    sql.NullString  `gorm:"column:contract_period" json:"contract_period"`
	ContractStartDate time.Time       `gorm:"column:contract_start_date" json:"contract_start_date"`
	ContractStopDate  time.Time       `gorm:"column:contract_stop_date" json:"contract_stop_date"`
	CreateTime        time.Time       `gorm:"column:create_time" json:"create_time"`
	DepartmentID      sql.NullInt64   `gorm:"column:department_id" json:"department_id"`
	Email             sql.NullString  `gorm:"column:email" json:"email"`
	EmergencyContact  sql.NullString  `gorm:"column:emergency_contact" json:"emergency_contact"`
	EmergencyPhone    sql.NullString  `gorm:"column:emergency_phone" json:"emergency_phone"`
	EntryID           sql.NullInt64   `gorm:"column:entry_id" json:"entry_id"`
	ForeignAbility    sql.NullString  `gorm:"column:foreign_ability" json:"foreign_ability"`
	FullName          sql.NullString  `gorm:"column:full_name" json:"full_name"`
	HealthStatus      sql.NullString  `gorm:"column:health_status" json:"health_status"`
	Height            sql.NullFloat64 `gorm:"column:height" json:"height"`
	HighestEducation  sql.NullString  `gorm:"column:highest_education" json:"highest_education"`
	HkAddress         sql.NullString  `gorm:"column:hk_address" json:"hk_address"`
	HkType            sql.NullString  `gorm:"column:hk_type" json:"hk_type"`
	HomeAddress       sql.NullString  `gorm:"column:home_address" json:"home_address"`
	HomeCode          sql.NullString  `gorm:"column:home_code" json:"home_code"`
	IsNeitui          sql.NullString  `gorm:"column:is_neitui" json:"is_neitui"`
	JobStatus         sql.NullInt64   `gorm:"column:job_status" json:"job_status"`
	Leader            sql.NullInt64   `gorm:"column:leader" json:"leader"`
	LeaveDate         time.Time       `gorm:"column:leave_date" json:"leave_date"`
	MaritalStatus     sql.NullString  `gorm:"column:marital_status" json:"marital_status"`
	Nation            sql.NullString  `gorm:"column:nation" json:"nation"`
	NeituiID          sql.NullInt64   `gorm:"column:neitui_id" json:"neitui_id"`
	Password          sql.NullString  `gorm:"column:password" json:"password"`
	PhoneNumber       sql.NullString  `gorm:"column:phone_number" json:"phone_number"`
	PhotoID           sql.NullInt64   `gorm:"column:photo_id" json:"photo_id"`
	PoliticsStatus    sql.NullString  `gorm:"column:politics_status" json:"politics_status"`
	PositionID        sql.NullInt64   `gorm:"column:position_id" json:"position_id"`
	PositiveDate      time.Time       `gorm:"column:positive_date" json:"positive_date"`
	PresentAddress    sql.NullString  `gorm:"column:present_address" json:"present_address"`
	PresentCode       sql.NullString  `gorm:"column:present_code" json:"present_code"`
	ProbationPeriod   sql.NullInt64   `gorm:"column:probation_period" json:"probation_period"`
	Profession        sql.NullString  `gorm:"column:profession" json:"profession"`
	Qq                sql.NullString  `gorm:"column:qq" json:"qq"`
	Relation          sql.NullString  `gorm:"column:relation" json:"relation"`
	Salt              sql.NullString  `gorm:"column:salt" json:"salt"`
	Sex               sql.NullInt64   `gorm:"column:sex" json:"sex"`
	Status            sql.NullInt64   `gorm:"column:status" json:"status"`
	UseMail           sql.NullInt64   `gorm:"column:use_mail" json:"use_mail"`
	UserID            int64           `gorm:"column:user_id;primary_key" json:"user_id"`
	UserName          sql.NullString  `gorm:"column:user_name" json:"user_name"`
	Weight            sql.NullFloat64 `gorm:"column:weight" json:"weight"`
	WorkNum           sql.NullInt64   `gorm:"column:work_num" json:"work_num"`
	YearsofWorking    sql.NullString  `gorm:"column:yearsof_working" json:"yearsof_working"`
}

// TableName sets the insert table name for this struct type
func (u *user) TableName() string {
	return "tb_global_user"
}
