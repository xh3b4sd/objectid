# objectid

Pseudo random numerical identifiers. The code below generates a numerical ID
consisting of the unix timestamp `1728336954` and the pseudo random number
`315444`. Why not just using UUIDs? If you can, you should. Purely numerical IDs
are particularly useful in sequential systems of contextually limited sets, on
which we want to apply set theoretical operations. A concrete example for this
would be object indexing in Redis sorted sets, where secondary keys are
numerical scores.

```
objectid.Random(objectid.Time(time.Now().UTC()))
```

```
1728336954315444
```
