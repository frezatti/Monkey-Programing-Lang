package ast

    type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expressions interface {
	Node
	expressionsNode()
}
