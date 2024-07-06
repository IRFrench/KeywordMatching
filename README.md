# Keyword matching

## The idea

Match keywords with wildcard characters. The idea is to be able to match the word `c*t` to `cat`.

Obviously the first thought would be regex, so why is this even a question? Its because regular expresions have a history of being slow, to the point that it can just lock up forever. We need to be able to match keywords quickly and efficiently without locking up the whole service.

This can be used for keyword filtering! Someone googles `p*rn` and they get... adult material! So we need them to... not... get that!

The actual keyword filtering itself is easy, if keyword in list then block. There are a bunch of considerations to make however, such as checking for wildcards in keywords. But there are more too, like spaces between letters, and all of these have to work together in a fraction of a second. This means that this matching function needs to work as quick as possible.

## The tests

Some pre-amble: These tests might only look for one wildcard character. This is for a couple of reasons. If we allow any amount of wildcards, then you can search for \*\*\* and it will find the word `cat`.
Also, since each wildcard check adds an exponential amount of time, a long list of wildcard characters would halt the service altogether, exactly why we are avoiding regex!
So to solve this, we need either need to stop the search early when it takes too long (not very useful), or limit the amount of wildcards we will search through. And for this I set the limit to one, since `c**` could be anything... like cat...

Anyways, onto the tests.

### Linear matching

Create all of the potential keywords, loop through the keyword list and see if the keywords are in there. Done.
This was more of a baseline to see how quick a simple check would be

### Regex matching

Turn the keyword into a regex pattern. Then check all of the keywords against the pattern.
This was also a baseline, helping to see how much we can improve on regex

### Rolling hash matching

So, this is a simple version of something that could be much more complex, but worked out well enough and quick enough that I don't think it needs much improvement

The idea is to "hash" the keyword list into a key value store. We asign each character an integer value, and use the sums of these values to make the map. Multiple words can have the same hash so we use this as a shortened list.

When we "hash" the keyword, we can compare it to the map and get a short list of possible keywords.

I have three implementations of this, custom mapping using multiplication, go rune mapping with multiplication and rune mapping with addition.

Runes in go are just letters mapped to numbers, so basically exactly what is needed. I don't think I have explained this very well but all the code is available if you wanna read it :)

## The results

For the keyword list, I used the list stored in `/usr/share/dict/words`. My laptop running Ubuntu desktop clocked in at 104,335 words.

I used both go benchmarking and custom stats to get the results. All of the implementations has zero allocations.

Three different keywords were used to test speeds with different keyword lengths. I also ran each test three times and took the median value, not as accurate as it could be but its good enough.

### Linear matching

| keyword | ns/op | avg time |
| --- | --- | --- |
| p\*rn | 0.002654 | 3.102312ms | 
| c\*t | 0.001664 | 1.840348ms |
| av\*cado | 0.005858 | 7.392907ms |

| stat name | stat |
| --- | --- |
| Fastest ns/op | 0.001646 |
| Slowest ns/op | 0.005919 |

### Regex matching

We have an extra test for regex, as this is able to handle more wildcards, so I was interested.

| keyword | ns/op | avg time |
| --- | --- | --- |
| p\*rn | 0.2703 | 283.377497ms | 
| c\*t | 0.2522 | 269.201659ms |
| av\*cado | 0.3566 | 369.329481ms |
| av\*c\*do  | 0.3535 | 376.817002ms |

| stat name | stat |
| --- | --- |
| Fastest ns/op | 0.2520 |
| Slowest ns/op | 0.3643 |

### Rolling hash matching - custom map

| keyword | ns/op | avg time
| --- | --- | --- |
| p\*rn | 0.0000260 | 14.827µs |
| c\*t | 0.0000233 | 20.175µs |
| av\*cado | 0.0000163 | 17.105µs |

*Segments are the list of keywords marked against a hash

| stat name | stat |
| --- | --- |
| Fastest ns/op | 0.0000151 |
| Slowest ns/op | 0.0000367 |
| max segment size | 68 |
| min segment size | 2 |
| avg segement size | 3 |
| median segement size | 2 |
| total segments | 33570 |

### Rolling hash matching - rune map - additive

| keyword | ns/op | avg time
| --- | --- | --- |
| p\*rn | 0.0000151 | 53.67µs |
| c\*t | 0.0000096 | 21.737µs |
| av\*cado | 0.0000203 | 63.625µs |

*Segments are the list of keywords marked against a hash

| stat name | stat |
| --- | --- |
| Fastest ns/op | 0.0000083 |
| Slowest ns/op | 0.0000509 |
| max segment size | 295 |
| min segment size | 2 |
| avg segement size | 60 |
| median segement size | 2 |
| total segments | 1734 |

### Rolling hash matching - rune map - multiplicative

| keyword | ns/op | avg time |
| --- | --- | --- |
| p\*rn | 0.0000124 | 9.35µs |
| c\*t | 0.0000147 | 9.362µs |
| av\*cado | 0.0000143 | 11.453µs |

*Segments are the list of keywords marked against a hash

| stat name | stat |
| --- | --- |
| Fastest ns/op | 0.0000121 |
| Slowest ns/op | 0.0000181 |
| max segment size | 9 |
| min segment size | 2 |
| avg segement size | 1 |
| median segement size | 2 |
| total segments | 94601 |
## Conclusions

Using a multiplicative hash map with go rune encoding worked really well, getting under the 10 nanosecond marker.

There might be other take aways to but eh, I just had fun making this :)
