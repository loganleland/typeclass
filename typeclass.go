package typeclass

//=============================================
// Eq: ==, !=
//=============================================

type Eq[T any] interface {
  Eq(T) bool
  Neq(T) bool
}

//=============================================
// Ordered: <, >, <=, >=, ==, !=
//=============================================

type Ord[T any] interface {
  Eq[T]
  Lesser(T) bool
  Greater(T) bool
  Leq(T) bool
  Geq(T) bool
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
