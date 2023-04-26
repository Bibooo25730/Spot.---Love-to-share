package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const sql = "root:11223344@tcp(localhost:3306)/panDemo?charset=utf8&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(sql))
	if err != nil {
		println(fmt.Errorf("%w", err))
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:       "gen/orm/dal",
		ModelPkgPath:  "gen/orm/model",
		Mode:          gen.WithDefaultQuery | gen.WithoutContext,
		FieldNullable: true,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)

	//g.ApplyBasic(g.GenerateModel("message", gen.FieldGenType("level", "level")))
	g.Execute()
}
