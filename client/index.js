import React from 'react';
import ReactDOM from 'react-dom';
import { WebSocketLink } from "@apollo/client/link/ws";
import {ApolloClient, InMemoryCache}  from '@apollo/client';
import { SubscriptionClient } from "subscriptions-transport-ws";
import { ApolloProvider } from '@apollo/react-hooks';
import App from './App';

const cache = new InMemoryCache();
const link = new WebSocketLink(
	new SubscriptionClient("ws://localhost:8080/query", {
		reconnect: true
	})
);
const apolloClient = new ApolloClient({
	link: link,
	cache: cache,
});


ReactDOM.render(
	<ApolloProvider client={apolloClient}>
		<App />
	</ApolloProvider>,
	document.getElementById('root')
);