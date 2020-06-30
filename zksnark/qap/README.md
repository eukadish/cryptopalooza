# Motivation

In my attempt to better understand Zero Knowledge Proof technology and spending too much time piecing together
information in random posts and articles. I decided to try get through some of the core work that has been built upon
too enable more sophisticated projects such as Zcash, recursive SNARKS (Coda Protocal), and Ethereum's scaling solution
with forgetful nodes.

So, I started with the scipr-labs repo: , and looking through the references found: as the meatiest paper to get through
for getting a feel for the nitty gritty of Zero Knowledge Proofs for algebraic expressions. The concept of an algebraic
expression is simple enough, so would be best suited to focus on the Zero Knowledgeness of the whole thing.

# Quadratic Arithmetic Programs

The main driver of the paper is to show how to take any arithmetic expression and convert it into a _Quadratic
Arithmetic Program_ from which is derivable a SNARK and then a statistically Zero Knowledge construction from the SNARK,
which could be non interactively verified.

So, to begin I took the simplest arithmetic expression I could think of: 2 x = 6 and moved forward with constructing a
_QAP_ for it. This is simple enough, but in order to create the 



In working through more advanced techniques in cryptography and blockchain applications understanding this construction
I found very important for more advanced protocols such as ZCash (which provides a great introduction to some of the key
mathematical insights) and the like.

However, when getting into some of the details involved with applying this technology many online resources mundge
together this fundamental work with more complicated techniques for the compiler construction and tolling built out by
scipr.

My goal here was to break out the work not worrying about efficiency or anything else within the construction and just
try nail down the concepts step by step along with more examples.

In addition I wanted to document here some of the more complex mathematical definitions for theunderlying concepts and
provide an easier to understand an intuitive feel to make it more approachable for the more day to day work.

h<sub>&theta;</sub>(x) = &theta;<sub>o</sub> x + &theta;<sub>1</sub>x

```math
SE = \frac{\sigma}{\sqrt{n}}
```

$`\sqrt{2}`$

## Definitions

I think something the paper lacks which confused me during it's initial interpretation is figuring out how to convert
a general arithmetic expression into one that is a product of expressions, so I look to provide better examples in the
sample functions here.

In addition the final construction for the ZK part is then fairly straightforward and generic to all such expressions.



## NIZK

Converting the QAP to a NIZK

## ZKSM

A protocol with some trusted setup to allow a trusted entity to prove
they are in possesion of a value without revealing it.
