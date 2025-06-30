import axios from "axios";

const apiUrl = "http://localhost:8000";

const Authorization = localStorage.getItem("token");

const Bearer = localStorage.getItem("token_type");


const requestOptions = {

  headers: {

    "Content-Type": "application/json",

    Authorization: `${Bearer} ${Authorization}`,

  },

};

async function ListExercises() {

  return await axios

    .get(`${apiUrl}/exercises`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

export { ListExercises };