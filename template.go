package main

type Package struct {
  Pkg string
  Service string
  Entity string
}

var code = `package {{ .Pkg }}

import (
  . "{{ .Service }}/domain/error"
  . "{{ .Service }}/domain/model"
  "sync"
)

type Handler func(Item) Item
type Filter func(Item) bool

type Functor interface {
  Map(Handler) Functor
}
type Monad interface {
  Fmap(Handler) Monad
  Filter(Filter) Monad
}

type {{ .Entity }}s struct {
  Value []{{ .Entity }}
  Error IError
}

func (items {{ .Entity }}s) Fmap(fn Handler) {{ .Entity }}s {
  if items.Error != nil {
    return items
  }
  ch := make(chan Item, len(items.Value))
  wg := new(sync.WaitGroup)
  wg.Add(len(items.Value))
  for _, v := range items.Value {
    go func(v {{ .Entity }}) {
      defer wg.Done()
      item := NewItem(v)
      ch <- fn(item)
    }(v)
  }
  wg.Wait()
  close(ch)
  
  result := make([]{{ .Entity }}, 0, len(items.Value))
  
  for item := range ch {
    switch {
    case item.Fail():
      return item.SetError(item.Error)
    default:
      result = append(result, item.Value)
    }
  }
  
  return items.SetValue(result...)
}
func (items {{ .Entity }}s) SetValue(market ...{{ .Entity }}) {{ .Entity }}s {
  values := make([]{{ .Entity }}, 0, len(market))
  values = append(values, market...)
  items.Value = values
  return items
}
func (items {{ .Entity }}s) Add(account {{ .Entity }}) {{ .Entity }}s {
  items.Value = append(items.Value, account)
  return items
}
func (items {{ .Entity }}s) SetError(err IError) {{ .Entity }}s {
  if err != nil {
    items.Error = err
  }
  return items
}

type Item struct {
  Value {{ .Entity }}
  Error IError
}

func (i Item) Fail() bool {
  if i.Error != nil {
    return true
  }
  return false
}
func (i Item) Success() bool {
  return !i.Fail()
}

type Items struct {
  Value []Item
  Error IError
}

func (i Items) Fail() bool {
  if i.Error != nil {
    return true
  }
  return false
}
func (i Items) Success() bool {
  if i.Fail() {
    return false
  }
  return true
}

func NewItem(key {{ .Entity }}) Item {
  return Item{
    Value: key,
    Error: nil,
  }
}
func (i Item) SetError(err IError) Item {
  i.Error = err
  return i
}

func (i Item) GetValue() {{ .Entity }} {
  return i.Value
}

func (i Item) Map(fn Handler) Item {
  return fn(i)
}

func (items {{ .Entity }}s) Fail() bool {
  if items.Error != nil {
    return false
  }
  return true
}
func (items {{ .Entity }}s) Success() bool {
  if items.Fail() {
    return false
  }
  return true
}
func (items {{ .Entity }}s) GetError() IError {
  return items.Error
}

func (items {{ .Entity }}s) GetValue() []{{ .Entity }} {
  return items.Value
}

func (items {{ .Entity }}s) Filter(fn Filter) {{ .Entity }}s {
  if items.Error != nil {
    return items
  }
  ch := make(chan Item, len(items.Value))
  wg := new(sync.WaitGroup)
  wg.Add(len(items.Value))
  for _, v := range items.Value {
    go func(v {{ .Entity }}) {
      defer wg.Done()
      item := NewItem(v)
      if fn(item) {
        ch <- item
      }
    }(v)
  }
  wg.Wait()
  close(ch)
  result := make([]{{ .Entity }}, 0, len(items.Value))
  for v := range ch {
    result = append(result, v.Value)
  }
  items.SetValue(result...)
  return items
}

`
