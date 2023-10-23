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
    grade := new(Grade)
    if err := c.Bind(grade); err != nil {
        return err
    }

    // 將成績紀錄添加到數據庫中
    _, err := Db.Exec("INSERT INTO grades (name, score) VALUES (?, ?)", grade.Name, grade.Score)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusCreated, grade)
}

// GetGrades 處理函數用於獲取所有成績紀錄
func GetGrades(c echo.Context) error {
    rows, err := Db.Query("SELECT name, score FROM grades")
    if err != nil {
        return err
    }
    defer rows.Close()

    grades := []Grade{}
    for rows.Next() {
        var grade Grade
        if err := rows.Scan(&grade.Name, &grade.Score); err != nil {
            return err
        }
        grades = append(grades, grade)
    }

    return c.JSON(http.StatusOK, grades)
}
