package typeclass

//=============================================
// Eq: ==, !=
//=============================================

type Eq[T any] interface {
  Eq(T) bool
}

func Neq[T Eq[T]](a, b T) bool {
  return !a.Eq(b)
}

//=============================================
// Ordered: <, >, <=, >=, ==, !=
//=============================================

type Ord[T any] interface {
  Eq[T]
  Lesser(T) bool
}

func Greater[T Ord[T]](a, b T) bool {
  return !a.Lesser(b) && !a.Eq(b)
}

func Leq[T Ord[T]](a, b T) bool {
  return a.Lesser(b) || a.Eq(b)
}

func Geq[T Ord[T]](a, b T) bool {
  return Greater(a, b) || a.Eq(b)
}

//=============================================
// Show
//=============================================

type Show[T any] interface {
  Show() string
}

//=============================================
// Bounded
//=============================================

type Bounded[T any] interface {
  Minbound() bool
  Maxbound() bool
}

//=============================================
// Functor: TBA
//
// fmap :: (a -> b) -> f a -> f b
//
// fmap id = id
// fmap (g . h) = (fmap g) . (fmap h)
//=============================================
