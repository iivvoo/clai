# CLAI

Command Line AI

A ChatGPT integration that helps you learn and write more complex shell commands. E.g.

```
$ clai Get third comma from tab-separated file and sort and count it
cut -f3 -d$'\t' yourfile.tsv | sort | uniq -c | sort -nr
(possibly some short explanation)

```


