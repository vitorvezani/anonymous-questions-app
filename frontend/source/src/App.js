import { useEffect, useState } from "react";

import Header from "./components/Header";

import styles from "./App.module.css";

const App = () => {
  const [questions, setQuestions] = useState([]);

  useEffect(() => {
    fetchQuestions();
  }, []);

  const fetchQuestions = async () => {
    try {
      const response = await fetch(`${process.env.REACT_APP_SERVER_URL}/api/v0/questions`);
      const questions = await response.json();
      setQuestions(questions);
    } catch (error) {
      if (error) {
        alert("Cannot fetch the list of questions. Please try again");
      }
    }
  };

  const upVote = (question) => async () => {
    if (!question) {
      return;
    }
    try {
      const response = await fetch(
        `${process.env.REACT_APP_SERVER_URL}/api/v0/questions/${question.ID}/up-vote`,
        {
          method: "POST",
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
          },
        }
      );
      const upvotedQuestion = await response.json();
      if (upvotedQuestion?.up_votes) {
        fetchQuestions();
      }
    } catch (error) {
      if (error) {
        alert("Cannot up vote the selected question. Please try again");
      }
    }
  };

  const deleteAllQuestions = async () => {
    try {
      await fetch(`${process.env.REACT_APP_SERVER_URL}/api/v0/questions`, {
        method: "DELETE",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
      });
      fetchQuestions();
    } catch (error) {
      if (error) {
        alert("Cannot delete all questions. Please try again");
      }
    }
  };

  return (
    <>
      <Header />
      <div className={styles.titleContainer}>
        {questions?.length ? (
          <div className={styles.titleSection}>
            <h3 styles={styles.title}>Questions</h3>
            <svg
              className={styles.delete}
              fill="none"
              onClick={deleteAllQuestions}
              stroke="currentColor"
              strokeWidth="1.5"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
          </div>
        ) : (
          <></>
        )}
      </div>
      <div className={styles.questionsContainer}>
        {questions?.length ? (
          questions.map((question) => (
            <div key={question.ID} className={styles.question}>
              <p className={styles.questionContent}>{question.Text}</p>
              <div className={styles.upVoteContainer}>
                <svg
                  className={styles.upVote}
                  fill="none"
                  onClick={upVote(question)}
                  stroke="currentColor"
                  strokeWidth="1.5"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M6.633 10.5c.806 0 1.533-.446 2.031-1.08a9.041 9.041 0 012.861-2.4c.723-.384 1.35-.956 1.653-1.715a4.498 4.498 0 00.322-1.672V3a.75.75 0 01.75-.75A2.25 2.25 0 0116.5 4.5c0 1.152-.26 2.243-.723 3.218-.266.558.107 1.282.725 1.282h3.126c1.026 0 1.945.694 2.054 1.715.045.422.068.85.068 1.285a11.95 11.95 0 01-2.649 7.521c-.388.482-.987.729-1.605.729H13.48c-.483 0-.964-.078-1.423-.23l-3.114-1.04a4.501 4.501 0 00-1.423-.23H5.904M14.25 9h2.25M5.904 18.75c.083.205.173.405.27.602.197.4-.078.898-.523.898h-.908c-.889 0-1.713-.518-1.972-1.368a12 12 0 01-.521-3.507c0-1.553.295-3.036.831-4.398C3.387 10.203 4.167 9.75 5 9.75h1.053c.472 0 .745.556.5.96a8.958 8.958 0 00-1.302 4.665c0 1.194.232 2.333.654 3.375z"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  />
                </svg>
                <span>{question.UpVotes}</span>
              </div>
            </div>
          ))
        ) : (
          <div className={styles.emptyQuestion}>
            We do not have any questions now!
          </div>
        )}
      </div>
    </>
  );
};

export default App;
