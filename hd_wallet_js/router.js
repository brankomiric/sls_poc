import { generate, restore } from "./routeHandlers.js";

export default (event) => {
  if (
    event.requestContext.http.path == "/wallet/generate" &&
    event.requestContext.http.method == "GET"
  ) {
    return generate(event.queryStringParameters);
  } else if (
    event.requestContext.http.path == "/wallet/restore" &&
    event.requestContext.http.method == "POST"
  ) {
    let body;
    try {
      body = JSON.parse(event.body);
    } catch {
      body = event.body;
    }
    return restore(body);
  } else {
    throw new Error(`${event.requestContext.http.path} match not found`);
  }
};
