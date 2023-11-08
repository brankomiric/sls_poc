"use strict";
import router from "./router.js";

export const wallet = async (event) => {
  try {
    const response = router(event);
    return {
      statusCode: 200,
      body: response,
    };
  } catch (err) {
    return {
      statusCode: 500,
      body: {
        message: err.toString(),
      },
    };
  }
};
