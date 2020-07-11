This project is for making proof of concept code samples to try out different crypto protocols for educational purposes.

  * Zero Knowledge
  * SNARKS
  * Homomorphic Encryption
  * . . . 

The terms and concepts in cryptography are endless as well as all the different relations between the concepts and how
different protocols or used to build applications. To try and resolve my confusion and aggregate more complicated cocepts
I had to scower the internet to understand Im using this repo for organization and refrence. For now it contains code for
a proof of concept as a n educational tool to check I understand some of the building blocks to more advanced protocols
like Zcash.

# Quadratic Arithmetic Programs

In an attempt to understand the technical details in the technology behind zcash and after going through several blog posts, videos, and
tutorials it finally got to the point where I had to go from the source papers and really understand the origins of how the protocol was
built up.

Moving forward I had to string together some blog posts with the papers along with the scipr labs repo and aside from a few additional
sources is fairly self contained in understanding how to construct snarks starting from an arithmetic circuit.

To make sure I understood how the whole thing worked and nail don a working example I went through the painful process of following
step by step the construction as described in the paper. I also hope that in the future it would be much easier to get a grasp
for some of this stuff by just using this one repo as a resource instead of having to scower the internet like I did.

But, be fair a lot more self contained projects now exist to understand everything here step by step.


This part of the construction



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
