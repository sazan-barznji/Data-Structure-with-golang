package main

type node struct {
    name    string
    ID      int
    address string
    child   [2]*node
    parent  *node
}

var root *node
var next *doublelink //this will be used the the converting to arr

func insertTree(t *node, name string, ID int, address string) *node {
    if t == nil {
        return &node{name, ID, address, nil, nil}
    } else {
        if address[0] < t.address[0] {
            t.child[0] = insertTree(t.child[0], name, ID, address)
        } else {
            t.child[1] = insertTree(t.child[1], name, ID, address)
        }
    }
    return t
}

