package main

import (
	"fmt"
	"log"

	"github.com/eugenekadish/cryptopalooza/zksnark/qap/examples"
)

// multiplication gate
// { x | a * x = b}

// f(x1) = a * x1 + b = 1 * (a * x1 + b) = p1(x1) * p2(x2)

// https://eprint.iacr.org/2012/215.pdf

// + Section 7.2 = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// f(x1) = 3 * x1
//       = (3) * (x1)
//       = p1(x1) * p2(x1)
//       = 6

// p1(x1) = c_{0} + Sigma_{i = 1}^{m - 1} c_{i} * x_{i}
//        = c_{0}
//        = 3

// p2(x1) = d_{0} + Sigma_{i = 1}^{m - 1} d_{i} * x_{i}
//        = d_{1} * x_{1}
//        = 1 * 2

// + Definition 11 - Qaudratic Arithmetic Program (QAP) Q:

// t(x) = x - r

// v_{0}(x) = c_{0} = 3
// v_{1}(x) = c_{1} = 0
// v_{2}(x)         = 0

// w_{0}(x) = d_{0} = 0
// w_{1}(x) = d_{1} = 1
// w_{2}(x)         = 0

// y_{0}(x)         = 0
// y_{1}(x)         = 0
// y_{2}(x)         = 1

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                               *
 *  3       x1  3        a1  3       2    p(a1) * p(a1) - a2 = (3) * (a1) - a2   *
 *   \     /     \     /      \     /                           = 3 * 2 - 6      *
 *    \   /       \   /        \   /                            = 0              *
 *     \ /         \ /          \ /                                              *
 *      *           *            *                                               *
 *      |           |            |                                               *
 *      |           |            |                                               *
 *      |           |            |                                               *
 *    f(x1)         a2           6                                               *
 *                                                                               *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// f(x1) = 2 * x1 + 5
//       = 1 * (5 + 2 * x1)
//       = p1(x1) * p2(x1)
//       = 11

// p1(x1) = c_{0} + Sigma_{i = 1}^{m - 1} c_{i} * x_{i}
//        = c_{0}
//        = 1

// p2(x1) = d_{0} + Sigma_{i = 1}^{m - 1} d_{i} * x_{i}
//        = d_{0} + d_{1} * x_{1}
//        = 5 + (2) * (3)

// + Definition 11 - Qaudratic Arithmetic Program (QAP) Q:

// t(x) = x - r

// v_{0}(x) = c_{0} = 1
// v_{1}(x) = c_{1} = 0
// v_{2}(x)         = 0

// w_{0}(x) = d_{0} = 5
// w_{1}(x) = d_{1} = 2
// w_{2}(x)         = 0

// y_{0}(x)         = 0
// y_{1}(x)         = 0
// y_{2}(x)         = 11

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                                                               *
 *          2       x1          2       a1            2       3    p(a1) * p(a1) - a2 = (1) * (5 + 2 * a1) - a2  *
 *           \     /             \     /               \     /                        = (1) * (5 + 2 * 3) - 11   *
 *            \   /               \   /                 \   /                         = 0                        *
 *             \ /                 \ /                   \ /                                                     *
 *      5       *           5       *             5       *                                                      *
 *       \     /             \     /               \     /                                                       *
 *        \   /               \   /                 \   /                                                        *
 *         \ /                 \ /                   \ /                                                         *
 *  1       +           1       +             1       +                                                          *
 *   \     /             \     /               \     /                                                           *
 *    \   /               \   /                 \   /                                                            *
 *     \ /                 \ /                   \ /                                                             *
 *      *                   *                     *                                                              *
 *      |                   |                     |                                                              *
 *      |                   |                     |                                                              *
 *      |                   |                     |                                                              *
 *    f(x1)                 a2                    11                                                             *
 *                                                                                                               *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// + Section 7.3 = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// f(x1, x2, x3, x4) = 4 * x1 * x2 - 7 * x2 + 3 * x3
//                   = x3 + 3 * x4 - 7 * x5
//                   = f2(x3, x4, x5)

// x2          => x5
// x3          => x4
// 4 * x1 * x2 => x3

// f1(x1, x2) = 4 * x1 * x2
//            = (4 * x1) * x2
//            = p1(x1, x2) * p2(x1, x2)
//            = 24

// p1(x1, x2) = c_{0} + Sigma_{i = 1}^{m - 1} c_{i} * x_{i}
//            = c_{0} + c_{1} * x_{1}
//            = 4 * 3

// p2(x1, x2) = d_{0} + Sigma_{i = 1}^{m - 1} d_{i} * x_{i}
//            = d_{0} + d_{1} * x_{1} + d_{2} * x_{2}
//            = 1 * 2

// + Definition 11 - Qaudratic Arithmetic Program (QAP) Q1: I_{1} = { 0, 1, 2, 3 }

// t1(x) = x - r1

// v1_{0}(x) = c_{0} = 0
// v1_{1}(x) = c_{1} = 4
// v1_{2}(x) = c_{2} = 0
// v1_{3}(x)         = 0

// w1_{0}(x) = d_{0} = 0
// w1_{1}(x) = d_{1} = 0
// w1_{2}(x) = d_{2} = 1
// w1_{3}(x)         = 0

// y1_{0}(x)         = 0
// y1_{1}(x)         = 0
// y1_{2}(x)         = 0
// y1_{3}(x)         = 1

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                                             *
 *  x1       x2  a1       a2  3       2    p1(a1, a2) * p2(a1, 2) - a3 = (4 * a1) * (a2) - a3  *
 *    \     /      \     /     \     /                                 = 4 * 3 * 2 - 24        *
 *   4 *   /      4 *   /     4 *   /                                  = 0                     *
 *      \ /          \ /         \ /                                                           *
 *       *            *           *                                                            *
 *       |            |           |                                                            *
 *       |            |           |                                                            *
 *       |            |           |                                                            *
 *   f(x1, x2)        a3          24                                                           *
 *                                                                                             *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// f2(x3, x4, x5) = x3 + 3 * x4 - 7 * x5
//                = 1 * (x3 + 3 * x4 - 7 * x5)
//                = p1(x3, x4, x5) * p2(x3, x4, x5)
//                = 24 + 3 * 1 - 7 * 2
//                = 13

// p1(x3, x4, x5) = c_{3} + Sigma_{i = 4}^{m - 1} c_{i} * x_{i}
//                = c_{3}
//                = 1

// p2(x3, x4, x5) = d_{3} + Sigma_{i = 4}^{m - 1} d_{i} * x_{i}
//                = d_{3} + d_{4} * x_{4} + d_{5} * x_{5}
//                = 24 + 3 * 1 - 7 * 2

// + Definition 11 - Qaudratic Arithmetic Program (QAP) Q2: I_{2} = { 3, 4, 5 }

// t2(x) = x - r2

// v1_{3}(x) = c_{3} = 1
// v1_{4}(x) = c_{4} = 0
// v1_{5}(x) = c_{5} = 0
// v1_{6}(x)         = 0

// w1_{3}(x) = d_{3} = 24
// w1_{4}(x) = d_{4} = 3
// w1_{5}(x) = d_{5} = -7
// w1_{6}(x)         = 0

// y1_{3}(x)         = 0
// y1_{4}(x)         = 0
// y1_{5}(x)         = 0
// y1_{6}(x)         = 1

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                                                                               *
 *        x3       x5      a3       a5     24       2    p(a3, a4, a5) * p(a3, a4, a5) - a6 = (1) * (a3 + 3 * a4 - 7 * a5) - a6  *
 *          \     /          \     /         \     /                                        = (1) * (24 + 3 * 1 - 7 * 2) - 13    *
 *           \   * 7          \   * 7         \   * 7                                       = 0                                  *
 *            \ /              \ /             \ /                                                                               *
 *    x4       -       a4       -       1       -                                                                                *
 *      \     /          \     /         \     /                                                                                 *
 *     3 *   /          3 *   /         3 *   /                                                                                  *
 *        \ /              \ /             \ /                                                                                   *
 *         +                +               +                                                                                    *
 *         |                |               |                                                                                    *
 *         * 1              * 1             * 1                                                                                  *
 *         |                |               |                                                                                    *
 *  f2(x3, x4, x5)          a6              13                                                                                   *
 *                                                                                                                               *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// Composed QAP

// t(x) = (x - r1) * (x - r2)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// f(x1, x2, x3, x4, x5) = x1 * (x2 - 2 * x3 ) * (x4 + 7 * x2)
//                       = x4 * (x5 + 7 * x6)
//                       = f2(x4, x5, x6)

// x2                 => x6
// x4                 => x5
// x1 * (x2 - 2 * x3) => x4

// f1(x1, x2, x3) = x1 * (x2 - 2 * x3)
//                = p1(x1, x2, x3) * p2(x1, x2, x3)
//                = 5 * (7 - 2 * 3)
//                = 5

// p1(x1, x2, x3) = c_{0} + Sigma_{i = 1}^{3} c_{i} * x_{i}
//                = c_{0} + c_{1} * x_{1}
//                = 1 * 5

// p2(x1, x2, x3) = d_{0} + Sigma_{i = 1}^{3} d_{i} * x_{i}
//                = d_{0} + d_{1} * x_{1} + d_{2} * x_{2} + d_{3} * x_{3}
//                = 1 * 7 - 2 * 3

// + Definition 11 - Qaudratic Arithmetic Program (QAP) Q1: I_{1} = { 0, 1, 2, 3, 4 }

// t1(x) = x - r1

// v1_{0}(x) = c_{0} = 0
// v1_{1}(x) = c_{1} = 1
// v1_{2}(x) = c_{2} = 0
// v1_{3}(x) = c_{3} = 0
// v1_{4}(x)         = 0

// w1_{0}(x) = d_{0} = 0
// w1_{1}(x) = d_{1} = 0
// w1_{2}(x) = d_{2} = 1
// w1_{3}(x) = d_{3} = 2
// v1_{4}(x)         = 0

// y1_{0}(x)         = 0
// y1_{1}(x)         = 0
// y1_{2}(x)         = 0
// y1_{3}(x)         = 0
// y1_{4}(x)         = 1

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                           #
#        x2       x3      a2      a3      7        3                                  p1(a1, a2, a3) * p2(a1, a2, a3) - a4 = (a1) * (a2 - 2 * a3) - a4 = 5  #
#          \     /         \     /         \      /                                                                        = (5) * (7 - 2 * 3) - 5          #
#           \   * 2         \   *           \    * 2                                                                       = 0                              #
#            \ /             \ /             \ /                                                                                                            #
#    x1       -      a1       -       5       -                                                                                                             #
#      \     /         \     /         \     /                                                                                                              #
#       \   /           \   /           \   /                                                                                                               #
#        \ /             \ /             \ /                                                                                                                #
#         *               *               *                                                                                                                 #
#         |               |               |                                                                                                                 #
#         |               |               |                                                                                                                 #
#         |               |               |                                                                                                                 #
#  f(x1, x2, x3)          a4              5                                                                                                                 #
#                                                                                                                                                           #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # */

// f2(x4, x5, x6) = x4 * (x5 + 7 * x6)
//                = p1(x4, x5, x6) * p2(x4, x5, x6)
//                = 5 * (2 + 7 * 1)
//                = 45

// p1(x4, x5, x6) = c_{3} + Sigma_{i = 4}^{6} c_{i} * x_{i}
//                = c_{3} + c_{4} * x_{4} + c_{5} * x_{5} + c_{6} * x_{6}
//                = c_{4} * x_{4}
//                = 1 * 5

// p2(x4, x5, x6) = d_{3} + Sigma_{i = 4}^{6} d_{i} * x_{i}
//                = d_{3} + d_{4} * x_{4} + d_{5} * x_{5} + d_{6} * x_{6}
//                = d_{5} * x_{5} + d_{6} * x_{6}
//                = 1 * 2 + 7 * 1

// + Definition 11 - Qaudratic Arithmetic Program (QAP) Q1: I_{2} = { 4, 5, 6, 7 }

// t2(x) = x - r2

// v1_{4}(x) = c_{4} = 1
// v1_{5}(x) = c_{5} = 0
// v1_{6}(x) = c_{6} = 0
// v1_{7}(x)         = 0

// w1_{4}(x) = d_{4} = 0
// w1_{5}(x) = d_{5} = 1
// w1_{6}(x) = d_{6} = 7
// v1_{7}(x)         = 0

// y1_{4}(x)         = 0
// y1_{5}(x)         = 0
// y1_{6}(x)         = 0
// y1_{7}(x)         = 1

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                           #
#        x5       x6      a5      a6      2        1                                      p1(a4, a5, a6) * p2(a4, a5, a6) - a7 = (a4) * (a5 + 7 * a6) - a7  #
#          \     /         \     /         \      /                                                                            = (5) * (2 + 7 * 1) - 45     #
#           \   * 7         \   * 7         \    * 7                                                                           = 0                          #
#            \ /             \ /             \ /                                                                                                            #
#    x4       +      a4       +       5       +                                                                                                             #
#      \     /         \     /         \     /                                                                                                              #
#       \   /           \   /           \   /                                                                                                               #
#        \ /             \ /             \ /                                                                                                                #
#         *               *               *                                                                                                                 #
#         |               |               |                                                                                                                 #
#         |               |               |                                                                                                                 #
#         |               |               |                                                                                                                 #
#  f(x4, x5, x6)          a7              45                                                                                                                #
#                                                                                                                                                           #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # */

// Composed QAP

// t(x) = (x - r1) * (x - r2)

// +
// +
// +

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// + Definition 11 - Qaudratic Arithmetic Program (QAP) Q1: I_{1} = { 0, 1, 2, 3 }

// t1(x) = x - r1

// v1_{0}(x) = c_{0} = 0
// v1_{1}(x) = c_{1} = 4
// v1_{2}(x) = c_{2} = 0
// v1_{3}(x)         = 0

// w1_{0}(x) = d_{0} = 0
// w1_{1}(x) = d_{1} = 0
// w1_{2}(x) = d_{2} = 1
// w1_{3}(x)         = 0

// y1_{0}(x)         = 0
// y1_{1}(x)         = 0
// y1_{2}(x)         = 0
// y1_{3}(x)         = 1

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                                             *
 *  x1       x2  a1       a2  3       2    p1(a1, a2) * p2(a1, 2) - a3 = (4 * a1) * (a2) - a3  *
 *    \     /      \     /     \     /                                 = 4 * 3 * 2 - 24        *
 *   4 *   /      4 *   /     4 *   /                                  = 0                     *
 *      \ /          \ /         \ /                                                           *
 *       *            *           *                                                            *
 *       |            |           |                                                            *
 *       |            |           |                                                            *
 *       |            |           |                                                            *
 *   f(x1, x2)        a3          24                                                           *
 *                                                                                             *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// f2(x1, x2, x3) = 7 * x1 * x2 - 7 * x2 + 3 * x3
//                = 7 * 3 *  + 5 - 4 * 3 * 2
//                = 2

// f2(x1, x2, ) = 7 * x1 + x2 - 4 * x1 * x3
//            = 7 * 3 + 5 - 4 * 3 * 2
//            = 2

// f1(x1, x2) = 4 * x1 * x2
//            = (4 * x1) * x2
//            = p1(x1, x2) * p2(x1, x2)
//            = 24

// 4 * x1 * x2 - 7 * x2 + 3 * x3
//                       = x3 + 3 * x4 - 7 * x5
//                       = f2(x3, x4, x5)

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                                                                                 *
 *         x3       x5      a3       a5     24       2    p(a3, a4, a5) * p(a3, a4, a5) - a6 = (1) * (a3 + 3 * a4 - 7 * a5) - a6   *
 *           \     /          \     /         \     /                                        = (1) * (24 + 3 * 1 - 7 * 2) - 13     *
 *            \   * 7          \   * 7         \   * 7                                       = 0                                   *
 *             \ /              \ /             \ /                                                                                *
 *     x4       -       a4       -       1       -                                                                                 *
 *       \     /          \     /         \     /                                                                                  *
 *      3 *   /          3 *   /         3 *   /                                                                                   *
 *         \ /              \ /             \ /                                                                                    *
 *          +                +               +                                                                                     *
 *          |                |               |                                                                                     *
 *          * 1              * 1             * 1                                                                                   *
 *          |                |               |                                                                                     *
 *   f2(x3, x4, x5)          a6              13                                                                                    *
 *                                                                                                                                 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                                           *
 *      x1      x1                                                                           *
 *       \     /                                                                             *
 *        \   /                                                                              *
 *         \ /                                                                               *
 *  a       *                          (r3) |                 (r3) |                 (r3) |  *
 *   \     /        - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  *
 *    \   /         | v1(r3) =          (1) | w1(r3) =         (0) | y1(r3) =         (0) |  *
 *     \ /          | v2(r3) =          (0) | w2(r3) =         (1) | y2(r3) =         (0) |  *
 *      *           | v3(r3) =          (0) | w3(r3) =         (0) | y3(r3) =         (1) |  *
 *      |           - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  *
 *      |           | v1(x)  =       x / r3 | w1(x)  = 0 || x - r3 | y1(x)  = 0 || x - r3 |  *
 *      |           | v2(x)  =  0 || x - r3 | w2(x)  =      x / r3 | y2(x)  = 0 || x - r3 |  *
 *      b           | v3(x)  =  0 || x - r3 | w3(x)  = 0 || x - r3 | y3(x)  =      x / r3 |  *
 *                  - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  *
 *                                                                                           *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// f(x1) = 4 * x1 * x1

// f(x1, x2) = 4 * x1^2 + 2 * x1 * x2 - 3 * x2^2
//           = 4 * x1 * x1 + 2 * x1 - 3 * x2 * x2

//

// now do with 1

// from the paper we have p1(X1) = x, p2(X1) = a

// linear equation
// { x | a x + b = 0 }

// the circuit is for the equation 1 * (a * x + b) = 0

// in terms of the QAP we have function F with inputs c_1 = 1, c_2 = -b / a, c_3 = 0

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*  1       a       x
*   \       \     /
*    \       \   /
*     \       \ /
*      \       *
*       \     /
*        \  +b
*         \ /
*          *
*          |
*          0
*
* = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
*
*       (r3) |       (r3) |       (r3) |
* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
* v1(r3) (1) | w1(r3) (0) | y1(r3) (0) |
* v2(r3) (0) | w2(r3) (1) | y2(r3) (0) |
* v3(r3) (0) | w3(r3) (0) | y3(r3) (1) |
*
*

// Now we need {vk(x)}, {wk(x)}, {yk(x)}

* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
* v1(x)  x / r3  | w1(x) x - r3 | y1(x) x - r3 |
* v2(x)  x - r3  | w2(x) x / r3 | y2(x) x - r3 |
* v3(x)  x - r3  | w3(x) x - r3 | y3(x) x / r3 |
*  * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// { x | a x^{2} + b x + c = 0 }

// a constant

// no multiply

// quadratic equation

// quadratic form

// pythagorean triples

// vitaik's example

// pinochio example

// paper example

// r1cs and not unit tests

// func calculate(a float64, b float64, c float64) (root float64, err error) {

// 	if a == 0 {
// 		err = fmt.Errorf("denominator variable a is 0")

// 		return root, err
// 	}

// 	if b*b < 4*a*c {
// 		err = fmt.Errorf("complex roots are not supported")

// 		return root, err
// 	}

// 	root = math.Sqrt(b*b - 4*a*c)

// 	return root, nil
// }

// func square(x int) {

// }

// func main() {

// 	// var vChal *big.Int
// 	// if vChal, err = rand.Int(
// 	// 	rand.Reader, bn256.Order,
// 	// ); err != nil {
// 	// 	e.Logger.
// 	// 		Error(err)

// 	// 	os.Exit(0)
// 	// }

// 	// var err error

// 	// var g *bn256.G1
// 	// if _, g, err = bn256.RandomG1(rand.Reader); err != nil {
// 	// 	log.Errorf("group element generation generation failed with error %v \n", err)

// 	// 	os.Exit(0)
// 	// }

// 	// var v_s
// 	// // var err error

// 	// var kp struct {
// 	// 	Pubk  *bn256.G1
// 	// 	Privk *big.Int
// 	// }

// 	// var gen *big.Int
// 	// gen, _ = new(big.Int).SetString(
// 	// 	"18560948149108576432482904553159745978835170526553990798435819795989606410925", 10,
// 	// )

// 	// var h = new(bn256.G2).ScalarBaseMult(gen)

// 	// kp.Privk, _ = rand.Int(rand.Reader, bn256.Order)
// 	// kp.Pubk = new(bn256.G1).ScalarBaseMult(kp.Privk)

// 	// var C = new(bn256.G2).ScalarBaseMult(x)
// 	// 	return C.Add(C, new(bn256.G2).ScalarMult(h, r))

// 	// if sigs[elem], err = bbsignatures.Sign(
// 	// 	// new(big.Int).SetInt64(int64(s[i])), v.kp.Privk,
// 	// 	new(big.Int).SetInt64(elem), v.kp.Privk,
// 	// ); err != nil {
// 	// 	return
// 	// }

// 	// Initialize variables

// 	// var m, _ = rand.Int(rand.Reader, bn256.Order)
// 	// var D = new(bn256.G2).ScalarMult(h, m)

// 	// // p.Commit, _ = new(bn256.G2).Unmarshal(commitBytes)

// 	// // D = g^s.H^m
// 	// var s, _ = rand.Int(rand.Reader, bn256.Order)
// 	// var aux = new(bn256.G2).ScalarBaseMult(s)
// 	// D.Add(D, aux)

// 	// var inv = new(big.Int).ModInverse(
// 	// 	new(big.Int).Mod(
// 	// 		new(big.Int).Add(m, kp.Privk), bn256.Order,
// 	// 	), bn256.Order,
// 	// )
// 	// var signature = new(bn256.G2).ScalarBaseMult(inv)

// 	// var tau, _ = rand.Int(rand.Reader, bn256.Order)
// 	// var proverSig = new(bn256.G2).ScalarMult(signature, tau)

// 	// var _, G1, _ = bn256.RandomG1(rand.Reader)
// 	// var _, G2, _ = bn256.RandomG2(rand.Reader)

// 	// // proof_out.V = new(bn256.G2).ScalarMult(proverSig, prover.tau)
// 	// var V = proverSig
// 	// var t, _ = rand.Int(rand.Reader, bn256.Order)
// 	// var a = bn256.Pair(G1, V)
// 	// a.ScalarMult(a, s)
// 	// // a.Invert(a)

// 	// var E = bn256.Pair(G1, G2)

// 	// a.Add(a, new(bn256.GT).ScalarMult(E, t))
// 	// D.Add(D, D)

// 	fmt.Printf("Linear QAP %t", examples.LinearQAP())

// 	fmt.Printf("hello, world\n")
// }

func main() {
	log.SetFlags(log.Lshortfile)

	// fmt.Printf("Linear QAP %t \n", examples.LinearQAP())
	fmt.Printf(" Example 2 %t \n", examples.Example2())
}
