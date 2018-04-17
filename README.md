## consul-kv-json

The Consul KV store has some nice tree properties for querying for a subtree (-recurse with a prefix). However, the output is still a flat list of KV pairs.

It would be nice if were possible to return a JSON representation of the subtree at a path. The inverse would also be nice, to convert a JSON blob into a list of KV pairs, accounting for key hierarchy.

This is a small `golang` implementation of that.