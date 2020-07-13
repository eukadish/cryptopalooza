Repo for educational purposes to build example constructions used in zero-knowledge protocols. The papers:

  * https://eprint.iacr.org/2012/215.pdf
  * https://eprint.iacr.org/2013/507.pdf

and these blog posts:

  * https://electriccoin.co/blog/snark-explain
  * https://medium.com/@VitalikButerin/quadratic-arithmetic-programs-from-zero-to-hero-f6d558cea649

describe how to generate a zero-knowledge proof for any arithmetic expression. This is done by creating first a SNARK
(Succinct Non-interactivity ARgument of Knowledge), which is derived from a QAP (Quadratic Arithmetic Program). The QAP
is the tricky part, so in each example is an arithmetic expression from which one is derived with the steps detailed in
the comments. They can also be used to cross reference with the documentation to make sure all the notation, indices,
polynomial interpolation, etc. is consistent.

To better understand the derivations on why QAPs and the example code work some basic facts about polynomial algebra are
useful to get a refresher on:

  * https://en.wikipedia.org/wiki/Lagrange_polynomial
  * https://en.wikipedia.org/wiki/Polynomial_remainder_theorem

Also, a simplifying technique of R1CS (Rank-1 Constraint Systems) to generate QAPs is shown in the comments and compared
with the derived QAP. The code verifies the QAP is correct by checking two sides of an equation with quadratic root
detection.

Also included is an example for zero-knowledge set membership:

  * https://www.ingwb.com/media/2667856/zero-knowledge-set-membership.pdf

More useful links on the topic:

  * https://github.com/scipr-lab/libsnark
  * https://github.com/matter-labs/awesome-zero-knowledge-proofs
