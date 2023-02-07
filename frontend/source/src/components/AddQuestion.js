import { useRef } from "react";

import Header from "./Header";

import styles from "./AddQuestion.module.css";

const AddQuestion = () => {
  const questionRef = useRef();

  const submit = async () => {
    const question = questionRef.current.value;
    questionRef.current.value = "";

    if (!question) {
      alert("Please input your question");
      return;
    }

    try {
      const response = await fetch(`${process.env.REACT_APP_SERVER_URL}/api/v0/questions`, {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ Text: question }),
      });
      const createdQuestion = await response.json();
      if (createdQuestion?.ID) {
        alert("Your question was added successfully");
      }
    } catch (error) {
      if (error) {
        console.log(error);
        alert("Cannot add your question. Please try again");
      }
    }
  };

  return (
    <>
      <Header />
      <div className={styles.container}>
        <input
          className={styles.input}
          placeholder="Enter Question..."
          ref={questionRef}
          type="text"
        />
        <button className={styles.submit} onClick={submit}>
          Submit
        </button>
      </div>
    </>
  );
};

export default AddQuestion;
