# Bender Hash

An Implementation of the Hash Table described in [Bender et al.](https://arxiv.org/abs/2111.00602)

## Usage

```golang

    d := bender.Hash{}

    err := d.Insert("some key", "some value")	
    value, err := d.Get("some key)
```
