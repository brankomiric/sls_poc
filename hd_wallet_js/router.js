import { generate, restore } from "./routeHandlers.js";

export default (event) => {
  if (event.path == "/wallet/generate" && event.httpMethod == "GET") {
    return generate(event.queryStringParameters);
  } else if (event.path == "/wallet/restore" && event.httpMethod == "POST") {
    return restore(event.body);
  } else {
    throw new Error("Invalid path or httpMethod");
  }
};
