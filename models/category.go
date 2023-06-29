package models

type Category struct {
  Id int
  Description string
}

func (s Store) GetCategories() ([]Category, error) {
  sql := "select * from category"
  rs, err := s.DB.Query(sql)
  if err != nil {
    return nil, err
  }
  defer rs.Close()
  var categories []Category
  for rs.Next() {
    var category Category
    err := rs.Scan(&category.Id, &category.Description)
    if err != nil {
      return nil, err
    }
    categories = append(categories, category)
  }
  return categories, nil
}

