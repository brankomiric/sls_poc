"use strict";
import router from "./router.js";

export const wallet = async (event) => {
  try {
    const response = router(event);
    console.log(response);
    return {
      statusCode: 200,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(response),
    };
  } catch (err) {
    console.log(err);
    return {
      statusCode: 500,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        message: err.toString(),
      }),
    };
  }
};
