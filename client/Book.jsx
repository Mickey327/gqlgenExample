import React from 'react';
import { useQuery } from '@apollo/react-hooks';
import * as GetBook from './GetBook.graphql';
import {gql} from "apollo-boost";
const GET_BOOK = gql`
query GetBook($id: ID!) {
	getBook(id: $id) {
		id
		title
		description
		author {
			id
			firstName
			lastName
		}
	}
}
`;
export default ({ id }) => {
	const { data, loading } = useQuery(GET_BOOK, {
		variables: {
			id: id
		}
	});

	const book = data ? data.getBook : null;

	return book ? (
		<div>
			<h1>{book.title}</h1>
		</div>
	) : (
		<div>Loading...</div>
	);
};