## consulkvjson 🌲

The Consul KV store has some nice tree properties for querying for a subtree (`-recurse` with a prefix). However, the output is still a flat list of KV pairs.

It would be nice if it were possible to return a JSON representation of the subtree at a path. The inverse would also be nice, to convert a JSON blob into a list of KV pairs, accounting for key hierarchy.

This is a small `golang` implementation of that.

Calling `ToKVs` with a JSON blob like this:
```js
{
  "key": {
    "at" : {
      "some": {
        "depth": true
      }
    }
  },
  "at_root": 123
}
```

Will return something like this:
```js
[
  {"key": "key/at/some/depth", "value": "true"},
  {"key": "at_root", "value": "123"},
]
```

And `ToJSON` will do the inverse. Note that JSON -> KVs turns all "leaves" of the JSON tree (numeric, boolean, string, null values) to strings.
