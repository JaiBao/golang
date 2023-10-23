// handlers.go

package api

import (

	"net/http"

	"github.com/labstack/echo/v4"
)

// Grade 結構用於儲存成績紀錄
type Grade struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// AddGrade 處理函數用於添加成績紀錄
func AddGrade(c echo.Context) error {
    // 解析請求中的 JSON 數據
    grade := new(Grade)
    if err := c.Bind(grade); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "無法解析 JSON 數據"})
    }

    // 將成績紀錄添加到數據庫
    _, err := db.Exec("INSERT INTO student (name, score) VALUES (?, ?)", grade.Name, grade.Score)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "無法插入數據"})
    }

    return c.JSON(http.StatusCreated, grade)
}

// GetGrades 處理函數用於獲取所有成績紀錄
func GetGrades(c echo.Context) error {
    // 查詢數據庫中的所有成績紀錄
    rows, err := db.Query("SELECT name, score FROM student")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "無法查詢數據庫"})
    }
    defer rows.Close()

    // 掃描數據庫中的結果
    var grades []Grade
    for rows.Next() {
        var grade Grade
        if err := rows.Scan(&grade.Name, &grade.Score); err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "無法掃描數據庫行"})
        }
        grades = append(grades, grade)
    }

    return c.JSON(http.StatusOK, grades)
}
