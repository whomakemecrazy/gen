package main

func New(p, s, e string) Package {
  return Package{
    Pkg: p,
    Service: s,
    Entity: e,
  }
}
